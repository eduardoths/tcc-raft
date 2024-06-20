package grpcutil

import (
	"context"

	pb "github.com/eduardoths/tcc-raft/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	Heartbeat(ctx context.Context, args *pb.HeartbeatArgs) (*pb.HeartbeatReply, error)
	RequestVote(ctx context.Context, args *pb.RequestVoteArgs) (*pb.RequestVoteReply, error)
	AppendEntries(ctx context.Context, args *pb.AppendEntriesArgs) (*pb.AppendEntriesReply, error)
}

func MakeClient(address string) Client {
	return makeGrpcClient(address)
}

type grpcClient struct {
	opts    []grpc.DialOption
	address string
	pb.UnimplementedRaftServer
}

func makeGrpcClient(address string) grpcClient {
	return grpcClient{
		opts: []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
		address: address,
	}
}

func (gc grpcClient) Heartbeat(ctx context.Context, args *pb.HeartbeatArgs) (*pb.HeartbeatReply, error) {
	conn, err := grpc.NewClient(gc.address, gc.opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewRaftClient(conn)
	return client.Heartbeat(ctx, args)
}

func (gc grpcClient) RequestVote(ctx context.Context, args *pb.RequestVoteArgs) (*pb.RequestVoteReply, error) {
	conn, err := grpc.NewClient(gc.address, gc.opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewRaftClient(conn)
	return client.RequestVote(ctx, args)
}

func (gc grpcClient) AppendEntries(ctx context.Context, args *pb.AppendEntriesArgs) (*pb.AppendEntriesReply, error) {
	conn, err := grpc.NewClient(gc.address, gc.opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewRaftClient(conn)
	return client.AppendEntries(ctx, args)
}
