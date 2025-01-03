package main

import (
	balancer "github.com/eduardoths/tcc-raft/cmd/balancer"
	"github.com/eduardoths/tcc-raft/cmd/node"
	"github.com/eduardoths/tcc-raft/cmd/orchestrator"
	restapi "github.com/eduardoths/tcc-raft/cmd/rest-api"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Raft database cli tool",
	Long:  "Single binary with multiple commands",
}

func init() {
	RootCmd.AddCommand(
		node.NodeCMD,
		orchestrator.OrchestratorCmd,
		balancer.BalancerCMD,
		restapi.RestApiCMD,
	)
}

func main() {
	Execute()
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}
