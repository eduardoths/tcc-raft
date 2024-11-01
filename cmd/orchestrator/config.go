package main

type Server struct {
	ID   string `yaml:"id"`
	Port string `yaml:"port"`
}

type RaftCluster struct {
	ServerCount       int      `yaml:"server_count"`
	Servers           []Server `yaml:"servers"`
	ElectionTimeout   int      `yaml:"election_timeout"`
	HeartbeatInterval int      `yaml:"heartbeat_interval"`
}

// Config is the root structure containing the Raft cluster configuration
type Config struct {
	RaftCluster RaftCluster `yaml:"raft_cluster"`
}
