package client

import (
	"context"

	"github.com/eduardoths/tcc-raft/dto"
	"github.com/eduardoths/tcc-raft/pkg/logger"
	pb "github.com/eduardoths/tcc-raft/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BalancerClient struct {
	conn   *grpc.ClientConn
	client pb.DatabaseClient
	logger logger.Logger
}

func NewBalancerClient(addr string, log logger.Logger) (*BalancerClient, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := pb.NewDatabaseClient(conn)

	return &BalancerClient{
		conn:   conn,
		client: client,
		logger: log,
	}, nil
}

func (bc *BalancerClient) Set(ctx context.Context, args dto.SetArgs) (dto.SetReply, error) {
	resp, err := bc.client.Set(ctx, args.ToProto())
	if err != nil {
		bc.logger.With("to-server", "balancer").Error(err, "Failed to create set request")
		return dto.SetReply{}, err
	}

	return dto.SetReplyFromProto(resp), nil
}

func (bc *BalancerClient) Get(ctx context.Context, args dto.GetArgs) (dto.GetReply, error) {
	resp, err := bc.client.Get(ctx, args.ToProto())
	if err != nil {
		bc.logger.With("to-server", "balancer").Error(err, "Failed to create get request")
		return dto.GetReply{}, err
	}

	return dto.GetReplyFromProto(resp), nil
}

func (bc *BalancerClient) Delete(ctx context.Context, args dto.DeleteArgs) (dto.DeleteReply, error) {
	resp, err := bc.client.Delete(ctx, args.ToProto())
	if err != nil {
		bc.logger.With("to-server", "balancer").Error(err, "Failed to create delete request")
		return dto.DeleteReply{}, err
	}

	return dto.DeleteReplyFromProto(resp), nil
}

func (bc *BalancerClient) GetLeader(ctx context.Context) (dto.GetLeaderReply, error) {
	resp, err := bc.client.GetLeader(ctx, &pb.EmptyDB{})
	if err != nil {
		bc.logger.With("to-server", "balancer").Error(err, "Failed to create get leader request")
		return dto.GetLeaderReply{}, err
	}
	return dto.GetLeaderReply{
		ID:   resp.GetId(),
		Term: int(resp.GetTerm()),
	}, nil
}
