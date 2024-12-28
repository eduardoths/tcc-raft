package raft

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/eduardoths/tcc-raft/internal/config"
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

	nodexMutex *sync.Mutex
	nodes      map[ID]*Node
	state      State

	// election data
	voteCount int

	// persistent stage on all servers
	electionMutex *sync.Mutex
	currentTerm   int
	votedFor      ID
	logEntry      []structs.LogEntry

	// volatile state on all servers
	commitIndex int
	lastApplied int

	// volatile state on leaders
	nextIdxMutex  *sync.Mutex
	nextIndex     map[ID]int
	matchIdxMutex *sync.Mutex
	matchIndex    map[ID]int

	//config
	heartbeatInterval int
	electionTimeout   int

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

func (r *Raft) getNodes() map[ID]*Node {
	r.nodexMutex.Lock()
	defer r.nodexMutex.Unlock()

	var cp = make(map[ID]*Node, len(r.nodes))

	for k, v := range r.nodes {
		cp[k] = v
	}

	return cp
}

func (r *Raft) setNodes(n map[ID]*Node) {
	r.nodexMutex.Lock()
	defer r.nodexMutex.Unlock()

	r.nodes = make(map[ID]*Node, len(n))
	for k, v := range n {
		r.nodes[k] = v
	}
}

func MakeRaft(id ID, nodes map[ID]*Node, log logger.Logger) *Raft {
	cfg := config.Get()
	r := &Raft{
		me:                id,
		nodexMutex:        &sync.Mutex{},
		nextIdxMutex:      &sync.Mutex{},
		matchIdxMutex:     &sync.Mutex{},
		electionMutex:     &sync.Mutex{},
		heartbeatInterval: cfg.RaftCluster.HeartbeatInterval,
		electionTimeout:   cfg.RaftCluster.ElectionTimeout,
	}
	r.setNodes(nodes)
	r.logger = log.With(
		"where", "raft",
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
		case <-time.After(time.Duration(r.electionTimeout+rand.Intn(150)) * time.Millisecond):
			r.logger.Info("Timed out, starting election")
			r.state = Candidate
		}
	case Candidate:
		r.logger.Debug("Starting election")
		r.currentTerm += 1
		r.votedFor = r.me
		r.voteCount = 1
		go r.broadcastRequestVote()

		select {
		case <-time.After(time.Duration(r.electionTimeout+rand.Intn(150)) * time.Millisecond):
			r.state = Follower
		case <-r.toLeaderC:
			r.state = Leader
			nodes := r.getNodes()
			r.nextIndex = make(map[ID]int, len(nodes))
			r.matchIndex = make(map[ID]int, len(nodes))
			for i := range nodes {
				r.setNextIndex(i, 1)
				r.setMatchIndex(i, 0)
			}
			r.logger.Info("Became a leader")
			r.persist()
			go func() {
				r.sendLeaderInfoToBalancer()
			}()

		}
	case Leader:
		r.broadcastHeartbeat()
		time.Sleep(time.Duration(r.heartbeatInterval) * time.Millisecond)
	}
}

func (r *Raft) getLastIndex() int {
	logLength := len(r.logEntry)
	if logLength == 0 {
		return 0
	}
	return r.logEntry[logLength-1].Index
}
