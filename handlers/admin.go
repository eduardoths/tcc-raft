package handlers

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/eduardoths/tcc-raft/dto"
	"github.com/eduardoths/tcc-raft/internal/config"
	"github.com/eduardoths/tcc-raft/pkg/client"
	"github.com/eduardoths/tcc-raft/pkg/logger"
)

type AdminHandler struct {
	mu      *sync.Mutex
	clients map[string]*client.AdminClient
	l       logger.Logger
}

func NewAdminHandler(addr string, servers map[string]string) *AdminHandler {
	ah := &AdminHandler{
		mu:      &sync.Mutex{},
		clients: make(map[string]*client.AdminClient),
		l:       logger.MakeLogger("handler", "admin", "server", "balancer"),
	}

	for id, v := range servers {
		client, err := client.NewAdminClient(v)
		if err != nil {
			ah.l.Error(err, "Failed to generate new admin client")
		}
		ah.clients[id] = client
	}

	return ah
}

func (ah *AdminHandler) AddNode(args dto.AddNodeArgs) error {
	cl, err := client.NewAdminClient(fmt.Sprintf("%s:%d", args.Host, args.Port))
	if err != nil {
		return err
	}
	ah.clients[args.ID] = cl

	ah.spawnNode(args)

	nodes := map[string]string{}
	for k, v := range ah.clients {
		nodes[k] = v.Addr
	}

	for _, cl := range ah.clients {
		if err := cl.SetNodes(context.Background(), nodes); err != nil {
			ah.l.Error(err, "failed to set nodes to server")
		}
	}

	config.AddNode(args.ID, args.Addr())

	return nil
}

func (ah *AdminHandler) spawnNode(args dto.AddNodeArgs) {
	ah.mu.Lock()
	defer ah.mu.Unlock()

	nodes := make(map[string]string)
	for k, v := range ah.clients {
		nodes[k] = v.Addr
	}

	nodesStr := config.Get().RaftCluster.NodesStr()
	nodesStr = fmt.Sprintf("%s,%s=%s", nodesStr, args.ID, args.Addr())

	execCmd := exec.Command(
		"./bin/cli", "grpc",
		"--id", args.ID,
		"--port", fmt.Sprintf("%d", args.Port),
		"--server_count", fmt.Sprintf("%d", len(ah.clients)+1),
		"--servers_map", nodesStr,
	)
	// Forward standard output and error
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
}

func (ah *AdminHandler) ShutdownNode(id string) error {
	delete(ah.clients, id)

	if err := ah.clients[id].Shutdown(context.Background()); err != nil {
		return err
	}

	nodes := map[string]string{}
	for k, v := range ah.clients {
		nodes[k] = v.Addr
	}

	for _, cl := range ah.clients {
		if err := cl.SetNodes(context.Background(), nodes); err != nil {
			ah.l.Error(err, "failed to set nodes to server")
		}
	}

	config.RemoveNode(id)

	return nil
}
