package raft

import (
	"context"

	"github.com/eduardoths/tcc-raft/dto"
	pb "github.com/eduardoths/tcc-raft/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	}

	return reply, nil
}

func (r *Raft) broadcastRequestVote() {
	args := dto.VoteArgs{
		Term:        r.currentTerm,
		CandidateID: r.me,
	}

	for i := range r.nodes {
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

	if r.voteCount >= len(r.nodes)/2+1 {
		r.toLeaderC <- true
	}

}

func (r *Raft) doRequestVote(serverID ID, args dto.VoteArgs) (dto.VoteReply, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(r.nodes[serverID].Address, opts...)
	if err != nil {
		return dto.VoteReply{}, err
	}
	defer conn.Close()
	client := pb.NewRaftClient(conn)
	serverReply, err := client.RequestVote(context.Background(), args.ToProto())
	if err != nil {
		return dto.VoteReply{}, err
	}
	return dto.VoteReplyFromProto(serverReply), nil
}
