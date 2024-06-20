package server

import (
	"context"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/eduardoths/tcc-raft/dto"
	"github.com/eduardoths/tcc-raft/pkg/logger"
	pb "github.com/eduardoths/tcc-raft/proto"
	"github.com/eduardoths/tcc-raft/raft"
	"google.golang.org/grpc"
)

type Server struct {
	mu     sync.Mutex
	raft   *raft.Raft
	logger logger.Logger

	grpc *grpc.Server

	pb.UnimplementedRaftServer
}

func CreateServer(raft *raft.Raft) *Server {
	return &Server{
		mu:     sync.Mutex{},
		raft:   raft,
		logger: logger.MakeLogger(),
	}
}

func (s *Server) Start(port string) {
	s.raft.Start()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		s.logger.Error(err, "failed to listen")
		os.Exit(1)
	}
	s.grpc = grpc.NewServer()
	pb.RegisterRaftServer(s.grpc, s)
	s.logger.Info("Listening at %v", lis.Addr())

	go func() {
		if err := s.grpc.Serve(lis); err != nil {
			s.logger.Error(err, "failed to serve")
			os.Exit(1)
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sc
	s.raft.Stop()
	s.grpc.GracefulStop()
}

func (s *Server) Heartbeat(ctx context.Context, args *pb.HeartbeatArgs) (*pb.HeartbeatReply, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	reply, err := s.raft.Heartbeat(ctx, dto.HeartbeatArgsFromProto(args))
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

	reply, err := s.raft.RequestVote(ctx, dto.VoteArgsFromProto(args))
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

func (s *Server) Set(ctx context.Context, args *pb.SetArgs) (*pb.SetReply, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	reply, err := s.raft.Set(ctx, dto.SetArgsFromProto(args))
	if err != nil {
		return &pb.SetReply{
			Index: -1,
			Noted: false,
		}, err
	}

	return reply.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, args *pb.DeleteArgs) (*pb.DeleteReply, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	reply, err := s.raft.Delete(ctx, dto.DeleteArgsFromProto(args))
	if err != nil {
		return &pb.DeleteReply{
			Index: -1,
			Noted: false,
		}, err
	}

	return reply.ToProto(), nil
}

func (s *Server) Get(ctx context.Context, args *pb.GetArgs) (*pb.GetReply, error) {
	reply, err := s.raft.Get(ctx, dto.GetArgsFromProto(args))
	if err != nil {
		return &pb.GetReply{
			Value: nil,
		}, err
	}

	return reply.ToProto(), nil
}

func (s *Server) SearchLog(ctx context.Context, args *pb.SearchLogArgs) (*pb.SearchLogReply, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	reply, err := s.raft.SearchLog(ctx, dto.SearchLogArgsFromProto(args))
	return reply.ToProto(), err
}
