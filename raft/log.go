package raft

import (
	"context"

	"github.com/eduardoths/tcc-raft/dto"
	"github.com/eduardoths/tcc-raft/structs"
)

func (r *Raft) SearchLog(ctx context.Context, args dto.SearchLogArgs) (dto.SearchLogReply, error) {
	if args.Index < 1 || r.getLastIndex() < args.Index {
		return dto.SearchLogReply{
			Command:  nil,
			Found:    false,
			Commited: false,
		}, nil
	}
	data := r.logEntry[args.Index-1]
	return dto.SearchLogReply{
		Command: &structs.LogCommand{
			Operation: data.Command.Operation,
			Key:       data.Command.Key,
			Value:     data.Command.Value,
		},
		Found:    true,
		Commited: r.commitIndex >= args.Index,
	}, nil
}
