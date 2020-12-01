package consensus

import (
	"context"
	"fmt"
	"go.uber.org/atomic"
	"log"
	"sync"
	"time"

	"github.com/joe-zxh/bycon/config"
	"github.com/joe-zxh/bycon/data"
	"github.com/joe-zxh/bycon/internal/logging"
	"github.com/joe-zxh/bycon/util"
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

	// from bycon
	Mut    sync.Mutex // Lock for all internal data
	ID     uint32
	TSeq   atomic.Uint32           // Total sequence number of next request
	seqmap map[data.EntryID]uint32 // Use to map {Cid,CSeq} to global sequence number for all prepared message
	View   uint32
	Apply  uint32 // Sequence number of last executed request

	Log    []*data.Entry
	LogMut sync.Mutex

	// Log        map[data.EntryID]*data.Entry // bycon的log是一个数组，因为需要保证连续，leader可以处理log inconsistency，而pbft不需要。client只有执行完上一条指令后，才会发送下一条请求，所以顺序 并没有问题。
	cps        map[int]*CheckPoint
	WaterLow   uint32
	WaterHigh  uint32
	F          uint32
	Q          uint32
	N          uint32
	monitor    bool
	Change     *time.Timer
	Changing   bool                // Indicate if this node is changing view
	state      interface{}         // Deterministic state machine's state
	ApplyQueue *util.PriorityQueue // 因为PBFT的特殊性(log是一个map，而不是list)，所以这里需要一个applyQueue。
	vcs        map[uint32][]*ViewChangeArgs
	lastcp     uint32

	Leader   uint32 // view改变的时候，再改变
	IsLeader bool   // view改变的时候，再改变

	waitEntry *sync.Cond
}

func (bycon *BYCONCore) AddCommand(command data.Command) {
	bycon.cmdCache.Add(command)
}

func (bycon *BYCONCore) CommandSetLen(command data.Command) int {
	return bycon.cmdCache.Len()
}

// CreateProposal creates a new proposal
func (bycon *BYCONCore) GetProposeCommands(timeout bool) *([]data.Command) {

	var batch []data.Command

	if timeout { // timeout的时候，不管够不够batch都要发起共识。
		batch = bycon.cmdCache.RetriveFirst(bycon.Config.BatchSize)
	} else {
		batch = bycon.cmdCache.RetriveExactlyFirst(bycon.Config.BatchSize)
	}

	return &batch
}

func (bycon *BYCONCore) InitLog() {
	bycon.LogMut.Lock()
	defer bycon.LogMut.Unlock()
	dPP := data.PrePrepareArgs{
		View:     1,
		Seq:      0,
		Commands: nil,
	}

	nonce := &data.Entry{
		PP:            &dPP,
		PreparedCert:  nil,
		Prepared:      true,
		CommittedCert: nil,
		Committed:     true,
		PreEntryHash:  &data.EntryHash{},
		Digest:        &data.EntryHash{},
		PrepareHash:   &data.EntryHash{},
		CommitHash:    &data.EntryHash{},
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

		// bycon
		ID:         uint32(conf.ID),
		seqmap:     make(map[data.EntryID]uint32),
		View:       1,
		Apply:      0,
		cps:        make(map[int]*CheckPoint),
		WaterLow:   0,
		WaterHigh:  2 * checkpointDiv,
		F:          uint32(len(conf.Replicas)-1) / 3,
		N:          uint32(len(conf.Replicas)),
		monitor:    false,
		Change:     nil,
		Changing:   false,
		state:      make([]interface{}, 1),
		ApplyQueue: util.NewPriorityQueue(),
		vcs:        make(map[uint32][]*ViewChangeArgs),
		lastcp:     0,
	}

	bycon.InitLog()

	bycon.waitEntry = sync.NewCond(&bycon.LogMut)

	bycon.Q = bycon.F*2 + 1
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
	ent.GetCommitHash() // 把Digest都算出来，以免后一个entry添加的时候 PreEntryHash是空的。记得在Put的外部，给ent加锁。
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
