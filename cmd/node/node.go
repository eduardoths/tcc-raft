package node

import (
	"fmt"
	"os"

	"github.com/eduardoths/tcc-raft/internal/config"
	"github.com/eduardoths/tcc-raft/pkg/logger"
	"github.com/eduardoths/tcc-raft/raft"
	"github.com/eduardoths/tcc-raft/server"
	"github.com/spf13/cobra"
)

var (
	NodeCMD = &cobra.Command{
		Use:   "node",
		Short: "Start new node containing a raft database server",
		Run:   startServer,
	}
	flags *config.Flags
)

func init() {
	flags = config.InitFlags(NodeCMD)

}

func startServer(cmd *cobra.Command, args []string) {
	log := logger.MakeLogger("cmd", "raft-database")
	config.InitWithCommands(flags, log)
	nodes := make(map[string]*raft.Node, 1)
	cfg := config.Get()
	for _, srv := range cfg.RaftCluster.Servers {
		nodes[srv.ID] = &raft.Node{
			Address: srv.Addr(),
		}
	}

	server := server.CreateServer(
		raft.MakeRaft(cfg.ID, nodes, cfg.Log),
		cfg.Log,
	)

	server.Start(fmt.Sprintf(":%d", cfg.Port))
}

func main() {
	if err := NodeCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
