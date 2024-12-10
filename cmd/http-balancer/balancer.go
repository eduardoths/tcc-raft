package httpbalancer

import (
	"github.com/eduardoths/tcc-raft/internal/config"
	"github.com/eduardoths/tcc-raft/pkg/logger"
	"github.com/eduardoths/tcc-raft/server"
	"github.com/spf13/cobra"
)

var (
	HttpBalancerCMD = &cobra.Command{
		Use:   "http-balancer",
		Short: "Start rest http api",
		Run:   startApi,
	}
	flags *config.Flags
)

func init() {
	flags = config.InitFlags(HttpBalancerCMD)
}

func startApi(cmd *cobra.Command, args []string) {
	log := logger.MakeLogger("cmd", "http-balancer")
	config.InitWithCommands(flags, log)

	s, err := server.NewHttpServer(log)
	if err != nil {
		panic(err)
	}

	s.Start()
}
