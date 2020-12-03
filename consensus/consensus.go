package consensus

import (
	"bytes"
	"context"
	"fmt"
	"github.com/joe-zxh/bycon/internal/proto"
	"go.uber.org/atomic"
	"log"
	"sync"
	"time"

	"github.com/joe-zxh/bycon/config"
	"github.com/joe-zxh/bycon/data"
	"github.com/joe-zxh/bycon/internal/logging"
)

const (
	changeViewTimeout = 60 * time.Second
	checkpointDiv     = 2000000
)

var logger *log.Logger

func init() {
	logger = logging.GetLogger()
}

// BYCONCore is the safety core of the BYCONCore protocol
type BYCONCore struct {
	// from hotstuff

	cmdCache *data.CommandSet // Contains the commands that are waiting to be proposed
	Config   *config.ReplicaConfig
	SigCache *data.SignatureCache
	cancel   context.CancelFunc // stops goroutines

	Exec chan []data.Command

	// from pbft
	Mut    sync.Mutex // Lock for all internal data
	ID     uint32
	TSeq   atomic.Uint32           // Total sequence number of next request
	seqmap map[data.EntryID]uint32 // Use to map {Cid,CSeq} to global sequence number for all prepared message
	View   uint32
	Apply  uint32 // Sequence number of last executed request

	Log                   []*data.Entry
	LogMut                sync.Mutex
	LastPreparedID        data.EntryID
	UpdateLastPreparedMut sync.Mutex

	cps       map[int]*CheckPoint
	WaterLow  uint32
	WaterHigh uint32
	F         uint32
	Q         uint32
	N         uint32
	monitor   bool
	Change    *time.Timer
	Changing  bool        // Indicate if this node is changing view
	state     interface{} // Deterministic state machine's state
	vcs       map[uint32][]*ViewChangeArgs
	lastcp    uint32

	Leader   uint32 // view改变的时候，再改变
	IsLeader bool   // view改变的时候，再改变

	waitEntry      *sync.Cond
	ViewChangeChan chan struct{}
	VoteFor        map[uint32]uint32              // key: view  value: nodeID
	VoteFinish     map[uint32]bool                // key: view value: is vote finish
	VoteCount      map[uint32](map[uint32]uint32) // key: view  value: (key: candidate value: count)
}

func (bycon *BYCONCore) AddCommand(command data.Command) {
	bycon.cmdCache.Add(command)
}

func (bycon *BYCONCore) CommandSetLen(command data.Command) int {
	return bycon.cmdCache.Len()
}

func (bycon *BYCONCore) GetProposeCommands(timeout bool) *([]data.Command) {

	var batch []data.Command

	if timeout { // timeout的时候，不管够不够batch都要发起共识。
		batch = bycon.cmdCache.RetriveFirst(bycon.Config.BatchSize)
	} else {
		batch = bycon.cmdCache.RetriveExactlyFirst(bycon.Config.BatchSize)
	}

	return &batch
}

// CreateProposal creates a new proposal
func (bycon *BYCONCore) CreateProposal(timeout bool) *data.PrePrepareArgs {

	var batch []data.Command

	if timeout { // timeout的时候，不管够不够batch都要发起共识。
		batch = bycon.cmdCache.RetriveFirst(bycon.Config.BatchSize)
	} else {
		batch = bycon.cmdCache.RetriveExactlyFirst(bycon.Config.BatchSize)
	}

	if batch == nil {
		return nil
	}
	e := &data.PrePrepareArgs{
		View:     bycon.View,
		Seq:      bycon.TSeq.Inc(),
		Commands: &batch,
	}
	return e
}

func (bycon *BYCONCore) InitLog() {
	bycon.LogMut.Lock()
	defer bycon.LogMut.Unlock()
	dPP := data.PrePrepareArgs{
		View:     0,
		Seq:      0,
		Commands: nil,
	}

	nonce := &data.Entry{
		PP:           &dPP,
		Committed:    true,
		PreEntryHash: &data.EntryHash{},
		Digest:       &data.EntryHash{},
	}

	bycon.Log = append([]*data.Entry{}, nonce)
}

// New creates a new BYCONCore instance
func New(conf *config.ReplicaConfig) *BYCONCore {
	logger.SetPrefix(fmt.Sprintf("hs(id %d): ", conf.ID))

	ctx, cancel := context.WithCancel(context.Background())

	bycon := &BYCONCore{
		// from hotstuff
		Config:   conf,
		cancel:   cancel,
		SigCache: data.NewSignatureCache(conf),
		cmdCache: data.NewCommandSet(),
		Exec:     make(chan []data.Command, 1),

		// from pbft
		ID:        uint32(conf.ID),
		seqmap:    make(map[data.EntryID]uint32),
		View:      1,
		Apply:     0,
		cps:       make(map[int]*CheckPoint),
		WaterLow:  0,
		WaterHigh: 2 * checkpointDiv,
		F:         uint32(len(conf.Replicas)-1) / 3,
		N:         uint32(len(conf.Replicas)),
		monitor:   false,
		Change:    nil,
		Changing:  false,
		state:     make([]interface{}, 1),
		vcs:       make(map[uint32][]*ViewChangeArgs),
		lastcp:    0,

		ViewChangeChan: make(chan struct{}, 1),
		VoteFor:        make(map[uint32]uint32),
		VoteFinish:     make(map[uint32]bool),
		VoteCount:      make(map[uint32](map[uint32]uint32)),
	}

	bycon.InitLog()

	bycon.waitEntry = sync.NewCond(&bycon.LogMut)

	bycon.Q = (bycon.F+bycon.N)/2 + 1
	bycon.Leader = (bycon.View-1)%bycon.N + 1
	bycon.IsLeader = (bycon.Leader == bycon.ID)

	// Put an initial stable checkpoint
	cp := bycon.getCheckPoint(-1)
	cp.Stable = true
	cp.State = bycon.state

	go bycon.proposeConstantly(ctx)

	return bycon
}

func (bycon *BYCONCore) proposeConstantly(ctx context.Context) {
	for {
		select {
		// todo: 一个计时器，如果是leader，就开始preprepare
		case <-ctx.Done():
			return
		}
	}
}

func (bycon *BYCONCore) Close() {
	bycon.cancel()
}

func (bycon *BYCONCore) GetExec() chan []data.Command {
	return bycon.Exec
}

func (bycon *BYCONCore) PutEntry(ent *data.Entry) {
	bycon.LogMut.Lock()
	defer bycon.LogMut.Unlock()

	listLen := uint32(len(bycon.Log))

	if listLen > ent.PP.Seq {
		panic(`try to overwrite an entry`)
	}

	for listLen < ent.PP.Seq {
		bycon.waitEntry.Wait()
		listLen = uint32(len(bycon.Log))
	}

	// listLen == ent.PP.Seq
	ent.PreEntryHash = bycon.Log[listLen-1].Digest
	ent.GetDigest()
	bycon.Log = append(bycon.Log, ent)

	bycon.waitEntry.Broadcast()
}

func (bycon *BYCONCore) GetEntryBySeq(seq uint32) *data.Entry {
	bycon.LogMut.Lock()
	defer bycon.LogMut.Unlock()

	listLen := uint32(len(bycon.Log))

	for listLen <= seq {
		bycon.waitEntry.Wait()
		listLen = uint32(len(bycon.Log))
	}

	return bycon.Log[seq]
}

func (bycon *BYCONCore) Prepared(ent *data.Entry) bool {
	if len(ent.P) >= int(bycon.Q) {
		// Key is the id of sender replica
		validSet := make(map[uint32]bool)
		for i, sz := 0, len(ent.P); i < sz; i++ {
			if ent.P[i].View == ent.PP.View && ent.P[i].Seq == ent.PP.Seq && bytes.Equal(ent.P[i].Digest[:], ent.Digest[:]) {
				validSet[ent.P[i].Sender] = true
			}
		}
		return len(validSet) > int(bycon.Q)
	}
	return false
}

// Locks : acquire s.lock before call this function
func (bycon *BYCONCore) Committed(ent *data.Entry) bool {
	if len(ent.C) > int(bycon.Q) {
		// Key is replica id
		validSet := make(map[uint32]bool)
		for i, sz := 0, len(ent.C); i < sz; i++ {
			if ent.C[i].View == ent.PP.View && ent.C[i].Seq == ent.PP.Seq && bytes.Equal(ent.C[i].Digest[:], ent.Digest[:]) {
				validSet[ent.C[i].Sender] = true
			}
		}
		return len(validSet) > int(bycon.Q)
	}
	return false
}

func (bycon *BYCONCore) ApplyCommands(commitSeq uint32) {
	bycon.Mut.Lock()
	defer bycon.Mut.Unlock()

	bycon.LogMut.Lock()
	defer bycon.LogMut.Unlock()

	for bycon.Apply < commitSeq {
		if bycon.Apply+1 < uint32(len(bycon.Log)) {
			ent := bycon.Log[bycon.Apply+1]
			bycon.Exec <- *ent.PP.Commands
			bycon.Apply++
		} else {
			break
		}
	}
}

func (bycon *BYCONCore) UpdateLastPreparedID(ent *data.Entry) {
	bycon.UpdateLastPreparedMut.Lock()
	defer bycon.UpdateLastPreparedMut.Unlock()

	if bycon.LastPreparedID.IsOlder(&data.EntryID{V: ent.PP.View, N: ent.PP.Seq}) {
		bycon.LastPreparedID.V = ent.PP.View
		bycon.LastPreparedID.N = ent.PP.Seq
	}
}

func (bycon *BYCONCore) UpdateVotes(pVC *proto.VoteConfirmArgs) {
	if !bycon.VoteFinish[pVC.NewView] {
		_, ok := bycon.VoteCount[pVC.NewView]
		if !ok {
			bycon.VoteCount[pVC.NewView] = make(map[uint32]uint32)
		}
		bycon.VoteCount[pVC.NewView][pVC.VoteFor]++

		if bycon.VoteCount[pVC.NewView][pVC.VoteFor] >= bycon.Q {

			logger.Printf("enter new view: %d\n", pVC.NewView)
			bycon.ViewChangeChan <- struct{}{}

			bycon.View = pVC.NewView
			bycon.Leader = pVC.VoteFor
			bycon.IsLeader = (bycon.ID == bycon.Leader)
			bycon.VoteFinish[pVC.NewView] = true
		}
	}
}
