package orchestrator

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/eduardoths/tcc-raft/internal/config"
	"github.com/eduardoths/tcc-raft/pkg/logger"
	"github.com/eduardoths/tcc-raft/raft"
	"github.com/spf13/cobra"
)

var (
	OrchestratorCmd = &cobra.Command{
		Use:   "orchestrator",
		Short: "Orchestrates raft nodes",
		Run:   RunOrchestrator,
	}
	flags *config.Flags
)

func init() {
	flags = config.InitFlags(OrchestratorCmd)
}

func RunOrchestrator(cmd *cobra.Command, args []string) {
	log := logger.MakeLogger("cmd", "orchestrator")
	config.InitWithYaml(log)
	cfg := config.Get()
	cfg.Port = flags.Port
	config.Set(cfg)

	var wg sync.WaitGroup

	nodes := make(map[raft.ID]*raft.Node, cfg.RaftCluster.ServerCount)
	for _, v := range cfg.RaftCluster.Servers {
		nodes[v.ID] = &raft.Node{
			Address: v.Addr(),
		}
	}
	nodesStr := cfg.RaftCluster.NodesStr()

	for _, s := range cfg.RaftCluster.Servers {
		execCmd := exec.Command(
			"./bin/cli", "grpc",
			"--id", s.ID,
			"--port", s.Port,
			"--election_timeout", fmt.Sprintf("%d", cfg.RaftCluster.ElectionTimeout),
			"--heartbeat_interval", fmt.Sprintf("%d", cfg.RaftCluster.HeartbeatInterval),
			"--server_count", fmt.Sprintf("%d", cfg.RaftCluster.ServerCount),
			"--servers_map", nodesStr,
		)
		// Forward standard output and error
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr

		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := execCmd.Run(); err != nil {
				log.Error(err, "Failed to run cmd")
			}
		}()
	}

	balancerCmd := exec.Command(
		"./bin/cli", "http-balancer",
		"--port", fmt.Sprintf("%d", cfg.Port),
		"--election_timeout", fmt.Sprintf("%d", cfg.RaftCluster.ElectionTimeout),
		"--heartbeat_interval", fmt.Sprintf("%d", cfg.RaftCluster.HeartbeatInterval),
		"--server_count", fmt.Sprintf("%d", cfg.RaftCluster.ServerCount),
		"--servers_map", nodesStr,
	)
	// Forward standard output and error
	balancerCmd.Stdout = os.Stdout
	balancerCmd.Stderr = os.Stderr

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := balancerCmd.Run(); err != nil {
			log.Error(err, "Failed to run balancer")
		}
	}()

	wg.Wait()
}
