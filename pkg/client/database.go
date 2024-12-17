package client

import (
	"context"

	"github.com/eduardoths/tcc-raft/dto"
	"github.com/eduardoths/tcc-raft/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DatabaseClient struct {
	client proto.DatabaseClient
	conn   *grpc.ClientConn
}

func NewDatabaseClient(addr string) (*DatabaseClient, error) {

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	client := proto.NewDatabaseClient(conn)
	return &DatabaseClient{
		conn:   conn,
		client: client,
	}, nil
}

func (dc *DatabaseClient) Close() error {
	return dc.conn.Close()
}

func (dc DatabaseClient) Set(ctx context.Context, args dto.SetArgs) (dto.SetReply, error) {
	rep, err := dc.client.Set(ctx, args.ToProto())
	if err != nil {
		return dto.SetReply{}, err
	}

	return dto.SetReplyFromProto(rep), nil
}

func (dc DatabaseClient) Get(ctx context.Context, args dto.GetArgs) (dto.GetReply, error) {
	resp, err := dc.client.Get(ctx, args.ToProto())
	if err != nil {
		return dto.GetReply{}, err
	}

	return dto.GetReplyFromProto(resp), nil
}

func (dc DatabaseClient) Delete(ctx context.Context, args dto.DeleteArgs) error {
	_, err := dc.client.Delete(ctx, args.ToProto())
	return err
}

func (dc DatabaseClient) GetLeader(ctx context.Context) (dto.SetLeader, error) {
	r, err := dc.client.GetLeader(ctx, &proto.EmptyDB{})

	if err != nil {
		return dto.SetLeader{}, err
	}

	return dto.SetLeader{
		ID:   r.GetId(),
		Term: int(r.GetTerm()),
	}, nil
}
