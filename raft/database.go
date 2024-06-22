package raft

import (
	"context"
	"errors"

	"github.com/eduardoths/tcc-raft/dto"
	"github.com/eduardoths/tcc-raft/storage"
	"github.com/eduardoths/tcc-raft/structs"
)

func (r *Raft) apply() error {
	if r.commitIndex < r.lastApplied {
		return errors.New("commit idx is lower than last applied")
	}
	for _, l := range r.logEntry {
		switch l.Command.Operation {
		case "SET":
			if err := r.storage.Save(storage.StorageSaveStruct{
				Key:   l.Command.Key,
				Value: l.Command.Value,
			}); err != nil {
				return err
			}
		case "DELETE":
			if err := r.storage.Delete(l.Command.Key); err != nil {
				return err
			}
		default:
			return errors.New("failed to apply invalid log entry")
		}
		r.lastApplied = l.Index
	}
	return nil
}

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

func (r *Raft) Delete(ctx context.Context, args dto.DeleteArgs) (dto.DeleteReply, error) {
	if r.state != Leader {
		return dto.DeleteReply{
			Index: -1,
			Noted: false,
		}, nil
	}
	idx := r.getLastIndex() + 1
	r.logEntry = append(r.logEntry, structs.LogEntry{
		Term:  r.currentTerm,
		Index: idx,
		Command: structs.LogCommand{
			Operation: "DELETE",
			Key:       args.Key,
		},
	})
	return dto.DeleteReply{
		Index: idx,
		Noted: true,
	}, nil
}

func (r *Raft) Get(ctx context.Context, args dto.GetArgs) (dto.GetReply, error) {
	b, err := r.storage.Get(args.Key)
	if err != nil {
		return dto.GetReply{
			Value: nil,
		}, err
	}
	return dto.GetReply{Value: b}, nil
}
