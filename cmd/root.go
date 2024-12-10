package main

import (
	"github.com/eduardoths/tcc-raft/cmd/grpc"
	httpbalancer "github.com/eduardoths/tcc-raft/cmd/http-balancer"
	"github.com/eduardoths/tcc-raft/cmd/orchestrator"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Raft database cli tool",
	Long:  "Single binary with multiple commands",
}

func init() {
	RootCmd.AddCommand(
		grpc.GrpcCMD,
		orchestrator.OrchestratorCmd,
		httpbalancer.HttpBalancerCMD,
	)
}

func main() {
	Execute()
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}
