package raft

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	pb "github.com/eduardoths/tcc-raft/proto"
	"google.golang.org/grpc"
)

type Server struct {
	mu   sync.Mutex
	raft *Raft

	grpc *grpc.Server

	pb.UnimplementedRaftServer
}

func CreateServer(raft *Raft) *Server {
	return &Server{
		mu:   sync.Mutex{},
		raft: raft,
	}
}

func (s *Server) Start(port string) {
	s.raft.start()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s.grpc = grpc.NewServer()
	pb.RegisterRaftServer(s.grpc, s)
	s.raft.logger.Info("Listening at %v", lis.Addr())

	go func() {
		if err := s.grpc.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sc
	s.grpc.GracefulStop()
}

func (s *Server) AppendEntries(ctx context.Context, args *pb.AppendEntriesArgs) (*pb.AppendEntriesReply, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.raft.AppendEntries(AppendEntriesArgs{}, &AppendEntriesReply{})

	return nil, err
}

func (s *Server) Heartbeat(ctx context.Context, args *pb.HeartbeatArgs) (*pb.HeartbeatReply, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	reply, err := s.raft.Heartbeat(ctx, HeartbeatArgsFromProto(args))
	if err != nil {
		return &pb.HeartbeatReply{
			Success:   false,
			Term:      -1,
			NextIndex: -1,
		}, err
	}
	return reply.ToProto(), nil
}

func (s *Server) RequestVote(ctx context.Context, args *pb.RequestVoteArgs) (*pb.RequestVoteReply, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	reply, err := s.raft.RequestVote(ctx, VoteArgsFromProto(args))
	if err != nil {
		reply.Term = -1
		reply.VoteGranted = false
		return &pb.RequestVoteReply{
			Term:        -1,
			VoteGranted: false,
		}, err
	}

	return reply.ToProto(), nil
}
