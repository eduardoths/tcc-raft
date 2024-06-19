package raft

type State int

const (
	Follower State = iota + 1
	Candidate
	Leader
)
