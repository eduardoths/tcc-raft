package raft

import (
	"encoding/gob"
	"os"

	"github.com/eduardoths/tcc-raft/structs"
)

type PersistenceData struct {
	State       State
	CurrentTerm int
	VotedFor    ID
	CommitIndex int
	LastApplied int
	NextIndex   map[ID]int
	MatchIndex  map[ID]int
	LogEntry    []structs.LogEntry
}

func (r *Raft) persist() {
	data := PersistenceData{
		State:       r.state,
		CurrentTerm: r.currentTerm,
		VotedFor:    r.votedFor,
		CommitIndex: r.commitIndex,
		LastApplied: r.lastApplied,
		NextIndex:   r.nextIndex,
		MatchIndex:  r.matchIndex,
		LogEntry:    r.logEntry,
	}

	r.logger.Debug("Persisting info")

	file, err := os.Create("data/persistency/" + r.me)
	if err != nil {
		r.logger.Error(err, "failed to create persistency file")
		panic(err)
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		r.logger.Error(err, "failed to create persistency file")
		panic(err)
	}
}

func (r *Raft) loadState() {
	file, err := os.Open("data/persistency/" + r.me)
	if err != nil {
		r.logger.Info("could not load previous state")
		return
	}
	defer file.Close()

	var data PersistenceData
	decoder := gob.NewDecoder(file)
	if err = decoder.Decode(&data); err != nil {
		r.logger.Error(err, "could not parse previous state, shutting down")
		panic("Could not decode previous state")
	}

	r.state = data.State
	r.currentTerm = data.CurrentTerm
	r.votedFor = data.VotedFor
	r.commitIndex = data.CommitIndex
	r.lastApplied = data.LastApplied
	r.matchIndex = data.MatchIndex
	r.nextIndex = data.NextIndex
	r.logEntry = data.LogEntry
	r.logger.Debug("finished recovering data from persistency files")
	r.logger.Info("loaded raft server")
}
