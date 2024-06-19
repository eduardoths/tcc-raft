package raft

import "github.com/eduardoths/tcc-raft/dto"

func (r *Raft) AppendEntries(args dto.AppendEntriesArgs) (dto.AppendEntriesReply, error) {
	return dto.AppendEntriesReply{}, nil
}
