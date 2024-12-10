package client

import (
	"context"

	"github.com/eduardoths/tcc-raft/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AdminClient struct {
	client proto.AdminClient
	conn   *grpc.ClientConn
	Addr   string
}

func NewAdminClient(addr string) (*AdminClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	client := proto.NewAdminClient(conn)
	return &AdminClient{
		conn:   conn,
		client: client,
		Addr:   addr,
	}, nil
}

func (dc *AdminClient) Close() error {
	return dc.conn.Close()
}

func (dc *AdminClient) SetNodes(ctx context.Context, nodes map[string]string) error {
	p := &proto.SetNodesArgs{
		Nodes: make(map[string]*proto.Node),
	}

	for k, v := range nodes {
		p.Nodes[k] = &proto.Node{
			Address: v,
		}
	}
	_, err := dc.client.SetNodes(ctx, p)
	return err
}

func (ac *AdminClient) Shutdown(ctx context.Context) error {
	_, err := ac.client.Shutdown(ctx, &proto.Empty{})
	return err
}
