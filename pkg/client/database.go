package client

import (
	"context"

	"github.com/eduardoths/tcc-raft/pkg/logger"
	pb "github.com/eduardoths/tcc-raft/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DatabaseClient struct {
	client pb.DatabaseClient
	conn   *grpc.ClientConn
	log    logger.Logger
}

func NewDatabaseClient(addr string, log logger.Logger) (*DatabaseClient, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	client := pb.NewDatabaseClient(conn)
	return &DatabaseClient{
		conn:   conn,
		client: client,
		log:    log,
	}, nil
}

func (dc *DatabaseClient) Close() error {
	return dc.conn.Close()
}

func (dc DatabaseClient) Set(ctx context.Context, args *pb.SetArgs) (*pb.SetReply, error) {
	resp, err := dc.client.Set(ctx, args)
	if err != nil {
		dc.log.Error(err, "failed to execute set request")
		return nil, err
	}

	return resp, nil
}

func (dc DatabaseClient) Get(ctx context.Context, args *pb.GetArgs) (*pb.GetReply, error) {
	resp, err := dc.client.Get(ctx, args)
	if err != nil {
		dc.log.Error(err, "failed to execute get request")
		return nil, err
	}
	return resp, nil
}

func (dc DatabaseClient) Delete(ctx context.Context, args *pb.DeleteArgs) (*pb.DeleteReply, error) {
	resp, err := dc.client.Delete(ctx, args)
	if err != nil {
		dc.log.Error(err, "failed to execute delete request")
		return nil, err
	}
	return resp, nil
}

func (dc DatabaseClient) GetLeader(ctx context.Context, args *pb.EmptyDB) (*pb.GetLeaderReply, error) {
	resp, err := dc.client.GetLeader(ctx, args)
	if err != nil {
		dc.log.Error(err, "failed to execute get leader request")
		return nil, err
	}

	return resp, nil
}
