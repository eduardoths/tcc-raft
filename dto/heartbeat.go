package dto

import (
	pb "github.com/eduardoths/tcc-raft/proto"
	"github.com/eduardoths/tcc-raft/structs"
)

type HeartbeatArgs struct {
	Term         int
	LeaderID     ID
	PrevLogIndex int
	PrevLogTerm  int
	Entries      []structs.LogEntry
	LeaderCommit int
}

func HeartbeatArgsFromProto(proto *pb.HeartbeatArgs) HeartbeatArgs {
	entries := make([]structs.LogEntry, len(proto.GetEntries()))
	for i, entry := range proto.GetEntries() {
		entries[i] = structs.LogEntry{
			Term:  int(entry.Term),
			Index: int(entry.Index),
			Command: structs.LogCommand{
				Operation: entry.Command.Operation,
				Key:       entry.Command.Key,
				Value:     entry.Command.Value,
			},
		}
	}
	return HeartbeatArgs{
		Term:         int(proto.GetTerm()),
		LeaderID:     proto.GetLeaderId(),
		PrevLogIndex: int(proto.GetPrevLogIndex()),
		PrevLogTerm:  int(proto.GetPrevLogTerm()),
		Entries:      entries,
		LeaderCommit: int(proto.GetLeaderCommit()),
	}
}

func (ha HeartbeatArgs) ToProto() *pb.HeartbeatArgs {
	entries := make([]*pb.LogEntry, len(ha.Entries))
	for i, entry := range ha.Entries {
		entries[i] = &pb.LogEntry{
			Term:  int32(entry.Term),
			Index: int32(entry.Index),
			Command: &pb.LogCMD{
				Operation: entry.Command.Operation,
				Key:       entry.Command.Key,
				Value:     entry.Command.Value,
			},
		}
	}
	return &pb.HeartbeatArgs{
		Term:         int32(ha.Term),
		LeaderId:     ha.LeaderID,
		PrevLogIndex: int32(ha.PrevLogIndex),
		PrevLogTerm:  int32(ha.PrevLogTerm),
		Entries:      entries,
		LeaderCommit: int32(ha.LeaderCommit),
	}
}

type HeartbeatReply struct {
	Success   bool
	Term      int
	NextIndex int
}

func HeartbeatReplyFromProto(proto *pb.HeartbeatReply) HeartbeatReply {
	return HeartbeatReply{
		Success:   proto.GetSuccess(),
		Term:      int(proto.GetTerm()),
		NextIndex: int(proto.GetNextIndex()),
	}
}

func (hr HeartbeatReply) ToProto() *pb.HeartbeatReply {
	return &pb.HeartbeatReply{
		Success:   hr.Success,
		Term:      int32(hr.Term),
		NextIndex: int32(hr.NextIndex),
	}
}
