package dto

import pb "github.com/eduardoths/tcc-raft/proto"

type SetBody struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (sb SetBody) ToArgs() SetArgs {
	return SetArgs{
		Key:   sb.Key,
		Value: []byte(sb.Value),
	}
}

type GetResponse struct {
	Value string `json:"string"`
}

type GetLeaderReply struct {
	ID   string `json:"string"`
	Term int    `json:"term"`
}

func GetLeaderReplyFromProto(proto *pb.GetLeaderReply) GetLeaderReply {
	return GetLeaderReply{
		ID:   proto.GetId(),
		Term: int(proto.GetTerm()),
	}
}
