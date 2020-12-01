package bycon

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/joe-zxh/bycon/data"
	"log"
	"net"
	"sync"
	"time"

	"github.com/joe-zxh/bycon/config"
	"github.com/joe-zxh/bycon/consensus"
	"github.com/joe-zxh/bycon/internal/logging"
	"github.com/joe-zxh/bycon/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var logger *log.Logger

func init() {
	logger = logging.GetLogger()
}

// BYCON is a thing
type BYCON struct {
	*consensus.BYCONCore
	proto.UnimplementedBYCONServer

	tls bool

	nodes map[config.ReplicaID]*proto.BYCONClient
	conns map[config.ReplicaID]*grpc.ClientConn

	server *byconServer

	closeOnce      sync.Once
	connectTimeout time.Duration
}

//New creates a new backend object.
func New(conf *config.ReplicaConfig, tls bool, connectTimeout, qcTimeout time.Duration) *BYCON {
	bycon := &BYCON{
		BYCONCore:      consensus.New(conf),
		nodes:          make(map[config.ReplicaID]*proto.BYCONClient),
		conns:          make(map[config.ReplicaID]*grpc.ClientConn),
		connectTimeout: connectTimeout,
	}
	return bycon
}

//Start starts the server and client
func (bycon *BYCON) Start() error {
	addr := bycon.Config.Replicas[bycon.Config.ID].Address
	err := bycon.startServer(addr)
	if err != nil {
		return fmt.Errorf("Failed to start GRPC Server: %w", err)
	}
	err = bycon.startClient(bycon.connectTimeout)
	if err != nil {
		return fmt.Errorf("Failed to start GRPC Clients: %w", err)
	}
	return nil
}

// 作为rpc的client端，调用其他hsserver的rpc。
func (bycon *BYCON) startClient(connectTimeout time.Duration) error {

	grpcOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithReturnConnectionError(),
	}

	if bycon.tls {
		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(bycon.Config.CertPool, "")))
	} else {
		grpcOpts = append(grpcOpts, grpc.WithInsecure())
	}

	for rid, replica := range bycon.Config.Replicas {
		if replica.ID != bycon.Config.ID {
			conn, err := grpc.Dial(replica.Address, grpcOpts...)
			if err != nil {
				log.Fatalf("connect error: %v", err)
				conn.Close()
			} else {
				bycon.conns[rid] = conn
				c := proto.NewBYCONClient(conn)
				bycon.nodes[rid] = &c
			}
		}
	}

	return nil
}

// startServer runs a new instance of byconServer
func (bycon *BYCON) startServer(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("Failed to listen to port %s: %w", port, err)
	}

	grpcServerOpts := []grpc.ServerOption{}

	if bycon.tls {
		grpcServerOpts = append(grpcServerOpts, grpc.Creds(credentials.NewServerTLSFromCert(bycon.Config.Cert)))
	}

	bycon.server = newBYCONServer(bycon)

	s := grpc.NewServer(grpcServerOpts...)
	proto.RegisterBYCONServer(s, bycon.server)

	go s.Serve(lis)
	return nil
}

// Close closes all connections made by the BYCON instance
func (bycon *BYCON) Close() {
	bycon.closeOnce.Do(func() {
		bycon.BYCONCore.Close()
		for _, conn := range bycon.conns { // close clients connections
			conn.Close()
		}
	})
}

// 这个server是面向 集群内部的。
type byconServer struct {
	*BYCON

	mut     sync.RWMutex
	clients map[context.Context]config.ReplicaID
}

func (bycon *BYCON) Propose(timeout bool) {
	dPP := bycon.CreateProposal(timeout)
	if dPP == nil {
		return
	}
	pPP := proto.PP2Proto(dPP)
	go bycon.BroadcastPrePrepareRequest(pPP)
	_, err := bycon.PrePrepare(context.TODO(), pPP)
	if err != nil {
		panic(err)
	}
}

func (bycon *BYCON) BroadcastPrePrepareRequest(pPP *proto.PrePrepareArgs) {
	logger.Printf("[B/PrePrepare]: view: %d, seq: %d, (%d commands)\n", pPP.View, pPP.Seq, len(pPP.Commands))

	for rid, client := range bycon.nodes {
		if rid != bycon.Config.ID {
			go func(id config.ReplicaID, cli *proto.BYCONClient) {
				_, err := (*cli).PrePrepare(context.TODO(), pPP)
				if err != nil {
					panic(err)
				}
			}(rid, client)
		}
	}
}

func (bycon *BYCON) BroadcastPrepareRequest(pP *proto.PrepareArgs) {
	logger.Printf("[B/Prepare]: view: %d, seq: %d\n", pP.View, pP.Seq)

	for rid, client := range bycon.nodes {

		if rid != bycon.Config.ID {
			go func(id config.ReplicaID, cli *proto.BYCONClient) {
				_, err := (*cli).Prepare(context.TODO(), pP)
				if err != nil {
					panic(err)
				}
			}(rid, client)
		}
	}
}

func (bycon *BYCON) BroadcastCommitRequest(pC *proto.CommitArgs) {
	logger.Printf("[B/Commit]: view: %d, seq: %d\n", pC.View, pC.Seq)

	for rid, client := range bycon.nodes {
		if rid != bycon.Config.ID {
			go func(id config.ReplicaID, cli *proto.BYCONClient) {
				_, err := (*cli).Commit(context.TODO(), pC)
				if err != nil {
					panic(err)
				}
			}(rid, client)
		}
	}
}

func (bycon *BYCON) PrePrepare(_ context.Context, pPP *proto.PrePrepareArgs) (*empty.Empty, error) {

	logger.Printf("PrePrepare: view:%d, seq:%d\n", pPP.View, pPP.Seq)

	dPP := pPP.Proto2PP()

	bycon.Mut.Lock()
	if !bycon.Changing && bycon.View == dPP.View {
		bycon.Mut.Unlock()

		ent := &data.Entry{
			PP: dPP,
		}
		ent.Mut.Lock()
		bycon.PutEntry(ent)

		pP := &proto.PrepareArgs{
			View:   dPP.View,
			Seq:    dPP.Seq,
			Digest: ent.GetDigest().ToSlice(),
			Sender: bycon.ID,
		}
		ent.Mut.Unlock()

		go bycon.BroadcastPrepareRequest(pP)
		go bycon.Prepare(context.TODO(), pP)

		return &empty.Empty{}, nil

	} else {
		bycon.Mut.Unlock()
		return &empty.Empty{}, errors.New(`正在view change 或者 view不匹配`)
	}
}

func (bycon *BYCON) Prepare(ctx context.Context, pP *proto.PrepareArgs) (*empty.Empty, error) {

	logger.Printf("Prepare: view:%d, seq:%d\n", pP.View, pP.Seq)

	dp := pP.Proto2P()
	bycon.Mut.Lock()

	if !bycon.Changing && bycon.View == dp.View {
		bycon.Mut.Unlock()
		ent := bycon.GetEntryBySeq(dp.Seq)
		ent.Mut.Lock()

		ent.P = append(ent.P, dp)
		if ent.PP != nil && !ent.SendCommit && bycon.Prepared(ent) {

			pC := &proto.CommitArgs{
				View:   ent.PP.View,
				Seq:    ent.PP.Seq,
				Digest: ent.Digest.ToSlice(),
				Sender: bycon.ID,
			}

			ent.SendCommit = true
			ent.Mut.Unlock()

			logger.Printf("[B/Commit]: view: %d, seq: %d\n", dp.View, dp.Seq)

			go bycon.BroadcastCommitRequest(pC)
			go bycon.Commit(context.TODO(), pC)
		} else {
			ent.Mut.Unlock()
		}
	} else {
		bycon.Mut.Unlock()
	}

	return &empty.Empty{}, nil
}

func (bycon *BYCON) Commit(_ context.Context, pC *proto.CommitArgs) (*empty.Empty, error) {

	logger.Printf("Commit: view:%d, seq:%d\n", pC.View, pC.Seq)

	dC := pC.Proto2C()

	bycon.Mut.Lock()

	if !bycon.Changing && bycon.View == dC.View {
		bycon.Mut.Unlock()
		ent := bycon.GetEntryBySeq(dC.Seq)

		ent.Mut.Lock()
		ent.C = append(ent.C, dC)
		if !ent.Committed && ent.SendCommit && bycon.Committed(ent) {
			logger.Printf("Committed entry: view: %d, seq: %d\n", ent.PP.View, ent.PP.Seq)
			ent.Committed = true
			ent.Mut.Unlock()

			go bycon.ApplyCommands(pC.Seq)
		} else {
			ent.Mut.Unlock()
		}
	} else {
		bycon.Mut.Unlock()
	}
	return &empty.Empty{}, nil
}

func newBYCONServer(bycon *BYCON) *byconServer {
	pbftSrv := &byconServer{
		BYCON:   bycon,
		clients: make(map[context.Context]config.ReplicaID),
	}
	return pbftSrv
}
