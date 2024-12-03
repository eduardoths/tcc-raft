package orchestrator

type Server struct {
	ID   string `yaml:"id" json:"id"`
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
}

type RaftCluster struct {
	ServerCount       int      `yaml:"server_count" json:"server_count"`
	Servers           []Server `yaml:"servers" json:"servers"`
	ElectionTimeout   int      `yaml:"election_timeout" json:"election_timeout"`
	HeartbeatInterval int      `yaml:"heartbeat_interval" json:"heartbeat_interval"`
}

type Config struct {
	RaftCluster RaftCluster `yaml:"raft_cluster" json:"raft_cluster"`
}
