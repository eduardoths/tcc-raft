package raft

import (
	"context"

	"github.com/eduardoths/tcc-raft/dto"
	grpcutil "github.com/eduardoths/tcc-raft/internal/util/grpc"
)

func (r *Raft) Heartbeat(ctx context.Context, args dto.HeartbeatArgs) (dto.HeartbeatReply, error) {
	if args.Term < r.currentTerm {
		return dto.HeartbeatReply{Success: false, Term: r.currentTerm, NextIndex: 0}, nil
	}

	if args.Term > r.currentTerm {
		func() {
			// keeping a separate function to unlock quickly
			r.electionMutex.Lock()
			defer r.electionMutex.Unlock()
			r.currentTerm = args.Term
			r.votedFor = args.LeaderID
			r.logger.Info("Updating current term")
		}()
		r.persist()
	}

	r.heartbeatC <- true
	if args.PrevLogIndex > r.getLastIndex() {
		return dto.HeartbeatReply{
			Success:   false,
			Term:      r.currentTerm,
			NextIndex: r.getLastIndex() + 1,
		}, nil
	}

	r.commitIndex = args.LeaderCommit
	r.logEntry = append(r.logEntry, args.Entries...)
	if len(args.Entries) == 0 {
		return dto.HeartbeatReply{Success: true, Term: r.currentTerm, NextIndex: 0}, nil
	}

	if err := r.apply(); err != nil {
		r.logger.Error(err, "shutting down server instance")
		panic(err)
	}

	return dto.HeartbeatReply{
		Success:   true,
		Term:      r.currentTerm,
		NextIndex: r.getLastIndex() + 1,
	}, nil
}

func (r *Raft) broadcastHeartbeat() {
	for i := range r.getNodes() {
		var args dto.HeartbeatArgs

		args.Term = r.currentTerm
		args.LeaderID = r.me
		args.LeaderCommit = r.commitIndex

		prevLogIndex := max(r.getNextIndex(i)-1, 0) // prevents errors if node is new

		if r.getLastIndex() > prevLogIndex {
			args.PrevLogIndex = prevLogIndex
			args.PrevLogTerm = r.logEntry[prevLogIndex].Term
			args.Entries = r.logEntry[prevLogIndex:]
		}

		go func(i ID, args dto.HeartbeatArgs) {
			r.sendHeartbeat(i, args)
		}(i, args)
	}
}

func (r *Raft) sendHeartbeat(serverID ID, args dto.HeartbeatArgs) {
	var reply dto.HeartbeatReply
	if serverID != r.me {
		var err error
		reply, err = r.doHeartbeat(serverID, args)
		if err != nil {
			reply.Success = false
			reply.Term = -1
		}
	} else {
		// simulated response from itself
		reply.Success = true
		reply.Term = r.currentTerm
		reply.NextIndex = r.getLastIndex() + 1
	}

	if reply.Success {
		if reply.NextIndex > 0 {
			r.setNextIndex(serverID, reply.NextIndex)
			r.setMatchIndex(serverID, r.getNextIndex(serverID)-1)
		}

		nextIdx := args.PrevLogIndex + len(args.Entries)

		if (nextIdx <= r.getLastIndex()) &&
			(r.commitIndex < nextIdx) &&
			r.logEntry[nextIdx-1].Term == r.currentTerm {
			count := 1
			for k := range r.getNodes() {
				if k != r.me && r.getMatchIndex(k) >= nextIdx {
					count += 1
				}
			}
			if count > len(r.getNodes())/2 {
				r.commitIndex = args.PrevLogIndex + len(args.Entries)
				if len(args.Entries) > 0 {
					r.persist()
				}
				r.apply()
			}
		}
	} else {
		if reply.Term > r.currentTerm {
			r.currentTerm = reply.Term
			r.state = Follower
			r.votedFor = ""
			r.persist()
			r.logger.Info("heartbeat term error")
		}
	}
}

func (r Raft) doHeartbeat(serverID ID, args dto.HeartbeatArgs) (dto.HeartbeatReply, error) {
	l := r.logger.With(
		"to", serverID,
		"leader-id", args.LeaderID,
		"term", args.Term,
	)

	l.Debug("Broadcasting heartbeat to server")
	response, err := grpcutil.MakeClient(r.getNodes()[serverID].Address).
		Heartbeat(context.Background(), args.ToProto())
	if err != nil {
		return dto.HeartbeatReply{}, err
	}
	l.With("success", response.GetSuccess()).Debug("Received response")
	return dto.HeartbeatReplyFromProto(response), nil
}
