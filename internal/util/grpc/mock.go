package grpcutil

import (
	"context"

	pb "github.com/eduardoths/tcc-raft/proto"
)

type mockController struct {
	heartbeat   func(ctx context.Context, args *pb.HeartbeatArgs) (*pb.HeartbeatReply, error)
	requestVote func(ctx context.Context, args *pb.RequestVoteArgs) (*pb.RequestVoteReply, error)
}

func (mc *mockController) SetHeartbeatResponse(reply *pb.HeartbeatReply, err error) {
	mc.heartbeat = func(ctx context.Context, args *pb.HeartbeatArgs) (*pb.HeartbeatReply, error) {
		return reply, err
	}
}

func (mc *mockController) SetRequestVoteResponse(reply *pb.RequestVoteReply, err error) {
	mc.requestVote = func(ctx context.Context, args *pb.RequestVoteArgs) (*pb.RequestVoteReply, error) {
		return reply, err
	}
}

var globalMock *mockController

func GetMockController() *mockController {
	if globalMock == nil {
		globalMock = new(mockController)
	}
	return globalMock
}

type mockClient struct {
	ctrl *mockController
}

func makeMockClient() mockClient {
	if globalMock == nil {
		globalMock = new(mockController)
	}

	return mockClient{
		ctrl: globalMock,
	}
}

func (gc mockClient) Heartbeat(ctx context.Context, args *pb.HeartbeatArgs) (*pb.HeartbeatReply, error) {
	return gc.ctrl.heartbeat(ctx, args)

}

func (gc mockClient) RequestVote(ctx context.Context, args *pb.RequestVoteArgs) (*pb.RequestVoteReply, error) {
	return gc.ctrl.requestVote(ctx, args)
}
