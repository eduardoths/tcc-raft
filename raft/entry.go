package raft

type AppendEntriesArgs struct {
}

type AppendEntriesReply struct{}

func (r *Raft) AppendEntries(args AppendEntriesArgs, reply *AppendEntriesReply) error {
	return nil
}
