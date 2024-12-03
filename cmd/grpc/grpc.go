package grpc

import (
	"fmt"
	"os"

	"github.com/eduardoths/tcc-raft/pkg/logger"
	"github.com/eduardoths/tcc-raft/raft"
	"github.com/eduardoths/tcc-raft/server"
	"github.com/spf13/cobra"
)

var (
	serverID          string
	log               logger.Logger
	port              int
	serverCount       int
	servers           map[string]string
	electionTimeout   int
	heartbeatInterval int
	enableK8s         bool
	GrpcCMD           = &cobra.Command{
		Use:   "grpc",
		Short: "Start grpc server",
		Run:   startServer,
	}
)

func init() {
	log = logger.MakeLogger("cmd", "grpc")
	// Define flags for the root command
	GrpcCMD.Flags().StringVar(&serverID, "id", "server-1", "Unique server ID")
	GrpcCMD.Flags().IntVar(&port, "port", 8080, "Port to run the HTTP server on")
	GrpcCMD.Flags().IntVar(&electionTimeout, "election_timeout", 5000, "Election timeout in milliseconds")
	GrpcCMD.Flags().IntVar(&heartbeatInterval, "heartbeat_interval", 1000, "Heartbeat interval in milliseconds")
	GrpcCMD.Flags().IntVar(&serverCount, "server_count", 1, "Number of servers running")
	GrpcCMD.Flags().BoolVar(&enableK8s, "enable_k8s", false, "Enable kubernetes support")
	GrpcCMD.Flags().StringToStringVar(
		&servers,
		"servers_map",
		map[string]string{serverID: fmt.Sprintf("[::]:%d", port)},
		"Map of nodes in the system",
	)

	permissions := os.FileMode(0755)
	if err := os.MkdirAll("data/persistency/", permissions); err != nil {
		panic(err)
	}
}

func startServer(cmd *cobra.Command, args []string) {
	nodes := make(map[string]*raft.Node, 1)
	for k, v := range servers {
		nodes[k] = &raft.Node{
			Address: v,
		}
	}

	server := server.CreateServer(
		raft.MakeRaft(serverID, nodes),
	)

	server.Start(fmt.Sprintf(":%d", port))
}

func main() {
	if err := GrpcCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
