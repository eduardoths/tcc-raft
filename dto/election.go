package dto

import (
	pb "github.com/eduardoths/tcc-raft/proto"
)

type ID = string

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
