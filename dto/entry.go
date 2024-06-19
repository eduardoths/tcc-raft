package dto

import "github.com/eduardoths/tcc-raft/structs"

type AppendEntriesArgs struct {
	Command structs.LogCommand
}

type AppendEntriesReply struct {
	Index    int
	Commited bool
}

type CommitEntryArgs struct{}

type CommitEntryReply struct{}

type ReceiveEntryArgs struct{}

type ReceiveEntryReply struct{}
