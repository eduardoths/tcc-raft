package raft

import (
	"math/rand"
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
	nextIndex  map[ID]int
	matchIndex map[ID]int

	// channels
	heartbeatC chan bool
	toLeaderC  chan bool
	stop       chan struct{}

	proto.UnimplementedRaftServer
}

func MakeRaft(id ID, nodes map[ID]*Node) *Raft {
	r := &Raft{
		me:    id,
		nodes: nodes,
	}
	r.logger = logger.MakeLogger(
		"server", id,
		"term", &r.currentTerm,
		"state", &r.state,
	)

	return r
}

func (r *Raft) Start() {
	// r.storage = storage.NewStorage(fmt.Sprintf("db-%s", r.me))
	r.state = Follower
	r.currentTerm = 0
	r.votedFor = ""
	r.heartbeatC = make(chan bool)
	r.toLeaderC = make(chan bool)
	r.stop = make(chan struct{})

	go func() {
		for {
			select {
			case <-r.stop:
				r.logger.Info("Stopping raft")
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
				r.nextIndex[i] = 1
				r.matchIndex[i] = 0
			}
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
