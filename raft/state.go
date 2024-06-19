package raft

type State int

const (
	Follower State = iota + 1
	Candidate
	Leader
)

func (s State) String() string {
	if s == Follower {
		return "Follower"
	}

	if s == Candidate {
		return "Candidate"
	}

	return "Leader"
}
