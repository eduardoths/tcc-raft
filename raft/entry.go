package raft

type AppendEntriesArgs struct {
	Command LogCommand
}

type AppendEntriesReply struct {
	Index    int
	Commited bool
}

func (r *Raft) AppendEntries(args AppendEntriesArgs) (AppendEntriesReply, error) {
	return AppendEntriesReply{}, nil
}

type CommitEntryArgs struct{}

type CommitEntryReply struct{}

type ReceiveEntryArgs struct{}

type ReceiveEntryReply struct{}
