package raft

import (
	"context"

	pb "github.com/eduardoths/tcc-raft/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HeartbeatArgs struct {
	Term         int
	LeaderID     ID
	PrevLogIndex int
	PrevLogTerm  int
	Entries      []LogEntry
	LeaderCommit int
}

func HeartbeatArgsFromProto(proto *pb.HeartbeatArgs) HeartbeatArgs {
	entries := make([]LogEntry, len(proto.GetEntries()))
	for i, entry := range proto.GetEntries() {
		entries[i] = LogEntry{
			Term:  int(entry.Term),
			Index: int(entry.Index),
			Command: LogCommand{
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

func (r *Raft) Heartbeat(ctx context.Context, args HeartbeatArgs) (HeartbeatReply, error) {
	if args.Term < r.currentTerm {
		return HeartbeatReply{false, r.currentTerm, 0}, nil
	}

	r.heartbeatC <- true
	if len(args.Entries) == 0 {
		return HeartbeatReply{true, r.currentTerm, 0}, nil
	}

	if args.PrevLogIndex > r.getLastIndex() {
		return HeartbeatReply{
			Success:   false,
			Term:      r.currentTerm,
			NextIndex: r.getLastIndex() + 1,
		}, nil
	}

	r.log = append(r.log, args.Entries...)
	r.commitIndex = r.getLastIndex()

	return HeartbeatReply{
		Success:   true,
		Term:      r.currentTerm,
		NextIndex: r.getLastIndex() + 1,
	}, nil
}

func (r *Raft) broadcastHeartbeat() {
	for i := range r.nodes {
		var args HeartbeatArgs

		args.Term = r.currentTerm
		args.LeaderID = r.me
		args.LeaderCommit = r.commitIndex

		prevLogIndex := r.nextIndex[i] - 1
		if r.getLastIndex() > prevLogIndex {
			args.PrevLogIndex = prevLogIndex
			args.PrevLogTerm = r.log[prevLogIndex].Term
			args.Entries = r.log[prevLogIndex:]
		}

		go func(i ID, args HeartbeatArgs) {
			r.sendHeartbeat(i, args)
		}(i, args)
	}
}

func (r *Raft) sendHeartbeat(serverID ID, args HeartbeatArgs) {
	var reply HeartbeatReply
	if serverID != r.me {
		var err error
		reply, err = r.doHeartbeat(serverID, args)
		if err != nil {
			reply.Success = false
			reply.Term = -1
		}
	}

	if reply.Success {
		if reply.NextIndex > 0 {
			r.nextIndex[serverID] = reply.NextIndex
			r.matchIndex[serverID] = r.nextIndex[serverID] - 1
		}
	} else {
		if reply.Term > r.currentTerm {
			r.currentTerm = reply.Term
			r.state = Follower
			r.votedFor = ""
		}
	}
}

func (r Raft) doHeartbeat(serverID ID, args HeartbeatArgs) (HeartbeatReply, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(r.nodes[serverID].Address, opts...)
	if err != nil {
		return HeartbeatReply{}, err
	}
	defer conn.Close()

	client := pb.NewRaftClient(conn)
	serverReply, err := client.Heartbeat(context.Background(), args.ToProto())
	if err != nil {
		return HeartbeatReply{}, err
	}
	return HeartbeatReplyFromProto(serverReply), nil
}
