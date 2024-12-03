package main

import (
	"github.com/eduardoths/tcc-raft/cmd/grpc"
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
	)
}

func main() {
	Execute()
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}
