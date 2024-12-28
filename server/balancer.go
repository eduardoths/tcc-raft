package server

import (
	"context"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/eduardoths/tcc-raft/internal/config"
	"github.com/eduardoths/tcc-raft/pkg/client"
	"github.com/eduardoths/tcc-raft/pkg/logger"
	pb "github.com/eduardoths/tcc-raft/proto"
	"google.golang.org/grpc"
)

type Balancer struct {
	mu    *sync.Mutex
	idxMu *sync.Mutex

	logger    logger.Logger
	dbClients map[string]*client.DatabaseClient
	leaderId  string
	ids       []string
	idx       int

	grpc *grpc.Server

	pb.UnimplementedDatabaseServer
}

func NewBalancer(log logger.Logger) *Balancer {
	balancer := &Balancer{
		mu:        &sync.Mutex{},
		idxMu:     &sync.Mutex{},
		logger:    log.With("where", "balancer"),
		dbClients: make(map[string]*client.DatabaseClient),
		ids:       make([]string, 0),
		idx:       -1,
	}

	cfg := config.Get()
	balancer.mu.Lock()
	defer balancer.mu.Unlock()
	for _, srv := range cfg.RaftCluster.Servers {
		client, err := client.NewDatabaseClient(srv.Addr(),
			balancer.logger.With("to-server", srv.ID),
		)
		if err != nil {
			balancer.logger.Error(err, "failed to set up create client")
			continue
		}
		balancer.dbClients[srv.ID] = client
		balancer.ids = append(balancer.ids, srv.ID)

	}
	return balancer
}

func (b *Balancer) next() (string, *client.DatabaseClient) {
	if len(b.ids) == 0 {
		return "", nil
	}

	b.idxMu.Lock()
	defer b.idxMu.Unlock()

	b.idx = (b.idx + 1) % len(b.ids)
	id := b.ids[b.idx]
	client := b.dbClients[id]
	return id, client
}

func (b *Balancer) getLeader(ctx context.Context) (string, *client.DatabaseClient, error) {
	var respLeader string
	if b.leaderId == "" {
		resp, err := b.GetLeader(ctx, &pb.EmptyDB{})
		if err != nil {
			b.logger.Error(err, "failed to set a leader")
			return "", nil, err
		}

		respLeader = resp.Id
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	if respLeader != "" && respLeader != b.leaderId {
		b.leaderId = respLeader
	}
	return b.leaderId, b.dbClients[b.leaderId], nil
}

func (b *Balancer) Start(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		b.logger.Error(err, "failed to listen")
		os.Exit(1)
	}

	b.grpc = grpc.NewServer()
	pb.RegisterDatabaseServer(b.grpc, b)

	b.logger.Info("Listening at %v", lis.Addr())

	go func() {
		if err := b.grpc.Serve(lis); err != nil {
			b.logger.Error(err, "failed to serve")
			os.Exit(1)
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	for _, c := range b.dbClients {
		c.Close()
	}
	b.grpc.GracefulStop()
}

func (b *Balancer) Get(ctx context.Context, args *pb.GetArgs) (*pb.GetReply, error) {
	_, client := b.next()

	return client.Get(ctx, args)
}

func (b *Balancer) Set(ctx context.Context, args *pb.SetArgs) (*pb.SetReply, error) {
	_, client, err := b.getLeader(ctx)
	if err != nil {
		return nil, err
	}

	return client.Set(ctx, args)
}

func (b *Balancer) Delete(ctx context.Context, args *pb.DeleteArgs) (*pb.DeleteReply, error) {
	_, client, err := b.getLeader(ctx)
	if err != nil {
		return nil, err
	}

	return client.Delete(ctx, args)
}

func (b *Balancer) GetLeader(ctx context.Context, args *pb.EmptyDB) (*pb.GetLeaderReply, error) {
	_, client := b.next()
	return client.GetLeader(ctx, args)
}
