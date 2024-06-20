package grpcutil

import (
	"context"

	pb "github.com/eduardoths/tcc-raft/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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
