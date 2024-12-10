package dto

import (
	pb "github.com/eduardoths/tcc-raft/proto"
	"github.com/eduardoths/tcc-raft/structs"
)

type SetArgs struct {
	Key   string
	Value []byte
}

func (sa SetArgs) ToProto() *pb.SetArgs {
	return &pb.SetArgs{
		Key:   sa.Key,
		Value: sa.Value,
	}
}

func SetArgsFromProto(proto *pb.SetArgs) SetArgs {
	return SetArgs{
		Key:   proto.GetKey(),
		Value: proto.GetValue(),
	}
}

type SetReply struct {
	Index int
	Noted bool
}

func SetReplyFromProto(proto *pb.SetReply) SetReply {
	return SetReply{
		Index: int(proto.GetIndex()),
		Noted: proto.GetNoted(),
	}
}

func (sr SetReply) ToProto() *pb.SetReply {
	return &pb.SetReply{
		Index: int32(sr.Index),
		Noted: sr.Noted,
	}
}

type DeleteArgs struct {
	Key string
}

func (da DeleteArgs) ToProto() *pb.DeleteArgs {
	return &pb.DeleteArgs{
		Key: da.Key,
	}
}

func DeleteArgsFromProto(proto *pb.DeleteArgs) DeleteArgs {
	return DeleteArgs{
		Key: proto.GetKey(),
	}
}

type DeleteReply struct {
	Index int
	Noted bool
}

func (dr DeleteReply) ToProto() *pb.DeleteReply {
	return &pb.DeleteReply{
		Index: int32(dr.Index),
		Noted: dr.Noted,
	}
}

type GetArgs struct {
	Key string
}

func (ga GetArgs) ToProto() *pb.GetArgs {
	return &pb.GetArgs{
		Key: ga.Key,
	}
}

func GetArgsFromProto(proto *pb.GetArgs) GetArgs {
	return GetArgs{
		Key: proto.GetKey(),
	}
}

type GetReply struct {
	Value []byte
}

func GetReplyFromProto(proto *pb.GetReply) GetReply {
	return GetReply{
		Value: proto.GetValue(),
	}
}

func (gr GetReply) ToProto() *pb.GetReply {
	return &pb.GetReply{
		Value: gr.Value,
	}
}

type SearchLogArgs struct {
	Index int
}

func SearchLogArgsFromProto(proto *pb.SearchLogArgs) SearchLogArgs {
	return SearchLogArgs{
		Index: int(proto.GetIndex()),
	}
}

type SearchLogReply struct {
	Command  *structs.LogCommand
	Found    bool
	Commited bool
}

func (sr SearchLogReply) ToProto() *pb.SearchLogReply {
	var command *pb.LogCMD
	if sr.Command != nil {
		command = &pb.LogCMD{
			Operation: sr.Command.Operation,
			Key:       sr.Command.Key,
			Value:     sr.Command.Value,
		}
	}
	return &pb.SearchLogReply{
		Command:  command,
		Found:    sr.Found,
		Commited: sr.Commited,
	}
}
