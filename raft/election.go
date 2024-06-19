package raft

import (
	"context"

	pb "github.com/eduardoths/tcc-raft/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type VoteArgs struct {
	Term        int
	CandidateID ID
}

func (va VoteArgs) ToProto() *pb.RequestVoteArgs {
	return &pb.RequestVoteArgs{
		Term:        int32(va.Term),
		CandidateId: va.CandidateID,
	}
}

func VoteArgsFromProto(proto *pb.RequestVoteArgs) VoteArgs {
	return VoteArgs{
		CandidateID: proto.GetCandidateId(),
		Term:        int(proto.GetTerm()),
	}
}

type VoteReply struct {
	Term        int
	VoteGranted bool
}

func (vr VoteReply) ToProto() *pb.RequestVoteReply {
	return &pb.RequestVoteReply{
		Term:        int32(vr.Term),
		VoteGranted: vr.VoteGranted,
	}
}

func VoteReplyFromProto(proto *pb.RequestVoteReply) VoteReply {
	return VoteReply{
		Term:        int(proto.GetTerm()),
		VoteGranted: proto.GetVoteGranted(),
	}
}

func (r *Raft) RequestVote(ctx context.Context, args VoteArgs) (VoteReply, error) {
	r.logger.Debug("Received vote request from server %s", args.CandidateID)
	reply := VoteReply{}
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
	args := VoteArgs{
		Term:        r.currentTerm,
		CandidateID: r.me,
	}

	for i := range r.nodes {
		go func(i ID) {
			r.sendRequestVote(i, args)
		}(i)
	}
}

func (r *Raft) sendRequestVote(serverID ID, args VoteArgs) {
	var reply VoteReply
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

func (r *Raft) doRequestVote(serverID ID, args VoteArgs) (VoteReply, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(r.nodes[serverID].Address, opts...)
	if err != nil {
		return VoteReply{}, err
	}
	defer conn.Close()
	client := pb.NewRaftClient(conn)
	serverReply, err := client.RequestVote(context.Background(), args.ToProto())
	if err != nil {
		return VoteReply{}, err
	}
	return VoteReplyFromProto(serverReply), nil
}
