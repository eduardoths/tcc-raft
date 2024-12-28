package balancer

import (
	"fmt"

	"github.com/eduardoths/tcc-raft/internal/config"
	"github.com/eduardoths/tcc-raft/pkg/logger"
	"github.com/eduardoths/tcc-raft/server"
	"github.com/spf13/cobra"
)

var (
	BalancerCMD = &cobra.Command{
		Use:   "balancer",
		Short: "Balance requests between servers",
		Run:   startBalancer,
	}
	flags *config.Flags
)

func init() {
	flags = config.InitFlags(BalancerCMD)
}

func startBalancer(cmd *cobra.Command, args []string) {
	log := logger.MakeLogger("cmd", "balancer")
	config.InitWithCommands(flags, log)

	s := server.NewBalancer(log)

	s.Start(fmt.Sprintf(":%d", config.Get().Port))
}
