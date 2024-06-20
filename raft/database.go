package raft

import (
	"context"

	"github.com/eduardoths/tcc-raft/dto"
	"github.com/eduardoths/tcc-raft/structs"
)

func (r *Raft) Set(ctx context.Context, args dto.SetArgs) (dto.SetReply, error) {
	if r.state != Leader {
		// temporarily avoid set on follower
		return dto.SetReply{
			Index: -1,
			Noted: false,
		}, nil

	}

	idx := r.getLastIndex() + 1
	r.logEntry = append(r.logEntry, structs.LogEntry{
		Term:  r.currentTerm,
		Index: idx,
		Command: structs.LogCommand{
			Operation: "SET",
			Key:       args.Key,
			Value:     args.Value,
		},
	})
	return dto.SetReply{
		Index: idx,
		Noted: true,
	}, nil

}
