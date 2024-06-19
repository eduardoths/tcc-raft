package raft

import (
	"context"
	"fmt"

	pb "github.com/eduardoths/tcc-raft/proto"
	"google.golang.org/grpc"
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
		fmt.Printf("Server %s (candidate) sending request to %s\n", r.me, serverID)
		var err error
		reply, err = r.doRequestVote(serverID, args)
		fmt.Printf("Server %s (candidate) got %t from %s\n", r.me, reply.VoteGranted, serverID)
		if err != nil {
			reply.VoteGranted = false
			reply.Term = -1
		}
	}

	if reply.Term > r.currentTerm {
		r.currentTerm = reply.Term
		r.state = Follower
		r.votedFor = ""
		return
	}

	if reply.VoteGranted {
		r.voteCount += 1
	}

	if r.voteCount >= len(r.nodes)/2+1 {
		r.toLeaderC <- true
	}

}

func (r *Raft) doRequestVote(serverID ID, args VoteArgs) (VoteReply, error) {
	opts := []grpc.DialOption{}
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
