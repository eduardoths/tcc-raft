package orchestrator

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/eduardoths/tcc-raft/pkg/logger"
	"github.com/eduardoths/tcc-raft/raft"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v3"
)

var (
	serverCount       int
	servers           map[string]string
	electionTimeout   int
	heartbeatInterval int
	enableK8s         bool
	log               logger.Logger
	OrchestratorCmd   = &cobra.Command{
		Use:   "orchestrator",
		Short: "Orchestrates raft nodes",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := initCfg()

			var wg sync.WaitGroup
			nodes := make(map[raft.ID]*raft.Node, serverCount)
			nodesStr := ""
			for _, v := range cfg.RaftCluster.Servers {
				addr := fmt.Sprintf("%s:%s", v.Host, v.Port)
				nodes[v.ID] = &raft.Node{
					Address: addr,
				}
				if nodesStr != "" {
					nodesStr = fmt.Sprintf("%s,", nodesStr)
				}
				nodesStr = fmt.Sprintf("%s%s=%s", nodesStr, v.ID, addr)
			}
			if strings.HasPrefix(",", nodesStr) {
				log.Info(nodesStr)
				nodesStr = nodesStr[1:]
			}

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
				go func(cmd *exec.Cmd) {
					defer wg.Done()
					if err := execCmd.Run(); err != nil {
						log.Error(err, "Failed to run cmd")
					}
				}(execCmd)

				wg.Wait()
			}
		},
	}
)

func init() {
	log = logger.MakeLogger("server", "orchestrator", "cmd", "orchestrator")
	OrchestratorCmd.Flags().IntVar(&electionTimeout, "election_timeout", 0, "Election timeout in milliseconds")
	OrchestratorCmd.Flags().IntVar(&heartbeatInterval, "heartbeat_interval", 0, "Heartbeat interval in milliseconds")
	OrchestratorCmd.Flags().IntVar(&serverCount, "server_count", 0, "Number of servers running")
	OrchestratorCmd.Flags().BoolVar(&enableK8s, "enable_k8s", false, "Enable kubernetes support")
	OrchestratorCmd.Flags().StringToStringVar(
		&servers,
		"servers_map",
		nil,
		"Map of nodes in the system",
	)

}

func initCfg() Config {
	cfg := *yamlConfig()
	if electionTimeout != 0 {
		cfg.RaftCluster.ElectionTimeout = electionTimeout
	}

	if heartbeatInterval != 0 {
		cfg.RaftCluster.HeartbeatInterval = heartbeatInterval
	}
	if serverCount != 0 {
		cfg.RaftCluster.ServerCount = serverCount
	}
	if servers != nil {
		cfg.RaftCluster.Servers = make([]Server, serverCount)
		for k, s := range servers {
			host, port, _ := net.SplitHostPort(s)
			cfg.RaftCluster.Servers = append(cfg.RaftCluster.Servers, Server{
				ID:   k,
				Host: host,
				Port: port,
			})
		}
	}

	permissions := os.FileMode(0755)
	if err := os.MkdirAll("data/persistency/", permissions); err != nil {
		log.Error(err, "failed to create directory")
		panic(err)
	}

	b, _ := json.Marshal(cfg)
	log.With("cfg", string(b)).Info("Finished loading config")

	return cfg
}

func yamlConfig() *Config {
	file, err := os.ReadFile("raft.yaml")
	if err != nil {
		log.With("err", err).Warn("Failed to open YAML file", "err", err)
		return nil
	}

	var config Config
	if err = yaml.Unmarshal(file, &config); err != nil {
		log.With("err", err).Info("Failed to parse YAML file", "err")
		return nil
	}

	return &config
}
