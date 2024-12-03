package raft

import (
	"context"

	"github.com/eduardoths/tcc-raft/dto"
	grpcutil "github.com/eduardoths/tcc-raft/internal/util/grpc"
)

func (r *Raft) RequestVote(ctx context.Context, args dto.VoteArgs) (dto.VoteReply, error) {
	r.logger.Debug("Received vote request from server %s", args.CandidateID)
	reply := dto.VoteReply{}
	if args.Term < r.currentTerm {
		reply.Term = r.currentTerm
		reply.VoteGranted = false
		return reply, nil
	}
	if r.votedFor == "" {
		r.currentTerm = args.Term
		r.votedFor = args.CandidateID
		reply.Term = r.currentTerm
		reply.VoteGranted = true
		r.persist()
	}

	return reply, nil
}

func (r *Raft) broadcastRequestVote() {
	args := dto.VoteArgs{
		Term:        r.currentTerm,
		CandidateID: r.me,
	}

	for i := range r.getNodes() {
		go func(i ID) {
			r.sendRequestVote(i, args)
		}(i)
	}
}

func (r *Raft) sendRequestVote(serverID ID, args dto.VoteArgs) {
	var reply dto.VoteReply
	if serverID != r.me {
		r.logger.Debug("Sending vote request to %s", serverID)
		var err error
		reply, err = r.doRequestVote(serverID, args)
		if err != nil {
			r.logger.Error(err, "failed to send request to server")
			reply.VoteGranted = false
			reply.Term = -1
		}

		r.logger.Debug("Received vote %t from server %s", reply.VoteGranted, serverID)

		if reply.Term > r.currentTerm {
			r.currentTerm = reply.Term
			r.state = Follower
			r.votedFor = ""
			return
		}

		if reply.VoteGranted {
			r.voteCount += 1
		}
	}

	if r.voteCount >= len(r.getNodes())/2+1 {
		r.toLeaderC <- true
	}

}

func (r *Raft) doRequestVote(serverID ID, args dto.VoteArgs) (dto.VoteReply, error) {
	response, err := grpcutil.MakeClient(r.getNodes()[serverID].Address).
		RequestVote(context.Background(), args.ToProto())
	if err != nil {
		return dto.VoteReply{}, err
	}
	return dto.VoteReplyFromProto(response), nil
}
