package raft

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/eduardoths/tcc-raft/pkg/logger"
	"github.com/eduardoths/tcc-raft/proto"
	"github.com/eduardoths/tcc-raft/storage"
	"github.com/eduardoths/tcc-raft/structs"
)

type ID = string

type Raft struct {
	logger  logger.Logger
	storage *storage.Storage
	me      ID

	nodes map[ID]*Node
	state State

	// election data
	voteCount int

	// persistent stage on all servers
	currentTerm int
	votedFor    ID
	logEntry    []structs.LogEntry

	// volatile state on all servers
	commitIndex int
	lastApplied int

	// volatile state on leaders
	nextIdxMutex  *sync.Mutex
	nextIndex     map[ID]int
	matchIdxMutex *sync.Mutex
	matchIndex    map[ID]int

	// channels
	heartbeatC chan bool
	toLeaderC  chan bool
	stop       chan struct{}

	proto.UnimplementedRaftServer
}

func (r *Raft) getNextIndex(id ID) int {
	r.nextIdxMutex.Lock()
	defer r.nextIdxMutex.Unlock()

	return r.nextIndex[id]
}

func (r *Raft) setNextIndex(id ID, v int) {
	r.nextIdxMutex.Lock()
	defer r.nextIdxMutex.Unlock()

	r.nextIndex[id] = v
}

func (r *Raft) getMatchIndex(id ID) int {
	r.matchIdxMutex.Lock()
	defer r.matchIdxMutex.Unlock()

	return r.matchIndex[id]
}

func (r *Raft) setMatchIndex(id ID, v int) {
	r.matchIdxMutex.Lock()
	defer r.matchIdxMutex.Unlock()

	r.matchIndex[id] = v
}

func MakeRaft(id ID, nodes map[ID]*Node) *Raft {
	r := &Raft{
		me:            id,
		nodes:         nodes,
		nextIdxMutex:  &sync.Mutex{},
		matchIdxMutex: &sync.Mutex{},
	}
	r.logger = logger.MakeLogger(
		"server", id,
		"term", &r.currentTerm,
		"state", &r.state,
	)

	return r
}

func (r *Raft) Start() {
	r.storage = storage.NewStorage(fmt.Sprintf("db-%s", r.me))
	r.state = Follower
	r.currentTerm = 0
	r.votedFor = ""
	r.heartbeatC = make(chan bool)
	r.toLeaderC = make(chan bool)
	r.stop = make(chan struct{})

	// tries to load state from file
	r.loadState()

	go func() {
		for {
			select {
			case <-r.stop:
				r.logger.Info("Stopping raft")
				r.storage.Shutdown()
				return
			default:
				r.mainLoop()
			}

		}
	}()
}

func (r *Raft) Stop() {
	r.stop <- struct{}{}
}

func (r *Raft) mainLoop() {
	switch r.state {
	case Follower:
		select {
		case <-r.heartbeatC:
			r.logger.Debug("Received heartbeat")
		case <-time.After(time.Duration(rand.Intn(150)+150) * time.Millisecond):
			r.logger.Debug("Timed out")
			r.state = Candidate
		}
	case Candidate:
		r.logger.Debug("Starting election")
		r.currentTerm += 1
		r.votedFor = r.me
		r.voteCount = 1
		go r.broadcastRequestVote()

		select {
		case <-time.After(time.Duration(rand.Intn(150)+150) * time.Millisecond):
			r.state = Follower
		case <-r.toLeaderC:
			r.state = Leader
			r.nextIndex = make(map[ID]int, len(r.nodes))
			r.matchIndex = make(map[ID]int, len(r.nodes))
			for i := range r.nodes {
				r.setNextIndex(i, 1)
				r.setMatchIndex(i, 0)
			}
			r.logger.Info("Became a leader")
			r.persist()
		}
	case Leader:
		r.logger.Debug("Sending heartbeat")
		r.broadcastHeartbeat()
		time.Sleep(50 * time.Millisecond)
	}
}

func (r *Raft) getLastIndex() int {
	logLength := len(r.logEntry)
	if logLength == 0 {
		return 0
	}
	return r.logEntry[logLength-1].Index
}
