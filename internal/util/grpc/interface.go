package grpcutil

import (
	"context"

	"github.com/caarlos0/env/v11"
	pb "github.com/eduardoths/tcc-raft/proto"
)

type Client interface {
	Heartbeat(ctx context.Context, args *pb.HeartbeatArgs) (*pb.HeartbeatReply, error)
	RequestVote(ctx context.Context, args *pb.RequestVoteArgs) (*pb.RequestVoteReply, error)
}

type config struct {
	MockGrpcClient bool `env:"MOCK_GRPC_CLIENT"`
}

var globalConfig config

func init() {
	if err := env.Parse(&globalConfig); err != nil {
		globalConfig.MockGrpcClient = false
	}
}

func MakeClient(address string) Client {
	if globalConfig.MockGrpcClient {
		return makeMockClient()
	}
	return makeGrpcClient(address)
}
