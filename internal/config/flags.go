package config

import (
	"github.com/spf13/cobra"
)

type Flags struct {
	ID                string
	Port              int
	ElectionTimeout   int
	HeartbeatInterval int
	ServerCount       int
	EnableK8s         bool
	Servers           map[string]string
}

func InitFlags(cmd *cobra.Command) *Flags {
	f := new(Flags)

	cmd.Flags().StringVar(&f.ID, "id", "", "Unique server ID")
	cmd.Flags().IntVar(&f.Port, "port", 0, "Port to run the HTTP server on")
	cmd.Flags().IntVar(&f.ElectionTimeout, "election_timeout", 0, "Election timeout in milliseconds")
	cmd.Flags().IntVar(&f.HeartbeatInterval, "heartbeat_interval", 0, "Heartbeat interval in milliseconds")
	cmd.Flags().IntVar(&f.ServerCount, "server_count", 0, "Number of servers running")
	cmd.Flags().BoolVar(&f.EnableK8s, "enable_k8s", false, "Enable kubernetes support")
	cmd.Flags().StringToStringVar(
		&f.Servers,
		"servers_map",
		make(map[string]string),
		"Map of nodes in the system",
	)
	return f
}
