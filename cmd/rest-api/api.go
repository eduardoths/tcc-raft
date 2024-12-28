package restapi

import (
	"github.com/eduardoths/tcc-raft/internal/config"
	"github.com/eduardoths/tcc-raft/pkg/logger"
	"github.com/eduardoths/tcc-raft/server"
	"github.com/spf13/cobra"
)

var (
	RestApiCMD = &cobra.Command{
		Use:   "rest-api",
		Short: "Start rest api server",
		Run:   startApi,
	}
	flags *config.Flags
)

func init() {
	flags = config.InitFlags(RestApiCMD)
}

func startApi(cmd *cobra.Command, args []string) {
	log := logger.MakeLogger("cmd", "http-balancer")
	config.InitWithCommands(flags, log)

	s, err := server.NewRestApiServer(log)
	if err != nil {
		panic(err)
	}

	s.Start()
}
