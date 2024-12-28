package config

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/eduardoths/tcc-raft/pkg/logger"
	"gopkg.in/yaml.v3"
)

var (
	globalConfig *Config
	mu           *sync.Mutex
)

type Server struct {
	ID   string `yaml:"id" json:"id"`
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
}

func (s Server) Addr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

type RaftCluster struct {
	ServerCount       int      `yaml:"server_count" json:"server_count"`
	Servers           []Server `yaml:"servers" json:"servers"`
	ElectionTimeout   int      `yaml:"election_timeout" json:"election_timeout"`
	HeartbeatInterval int      `yaml:"heartbeat_interval" json:"heartbeat_interval"`
	EnableK8s         bool     `yaml:"enable_k8s" json:"enable_k8s"`
	BalancerPort      int      `yaml:"balancer_port" json:"balancer_port"`
	BalancerHost      int      `yaml:"balancer_host" json:"balancer_host"`
	RestPort          int      `yaml:"rest_port" json:"rest_port"`
	RestHost          int      `yaml:"rest_host" json:"rest_host"`
}

func (rc RaftCluster) NodesStr() string {
	nodesStr := ""
	for _, v := range rc.Servers {
		addr := fmt.Sprintf("%s:%s", v.Host, v.Port)
		if nodesStr != "" {
			nodesStr = fmt.Sprintf("%s,", nodesStr)
		}
		nodesStr = fmt.Sprintf("%s%s=%s", nodesStr, v.ID, addr)
	}
	if strings.HasPrefix(",", nodesStr) {
		nodesStr = nodesStr[1:]
	}
	return nodesStr
}

type Config struct {
	RaftCluster RaftCluster `yaml:"raft_cluster" json:"raft_cluster"`
	ID          string      `yaml:"-" json:"id"`
	Port        int         `yaml:"-" json:"port"`
	Log         logger.Logger
}

func init() {
	if globalConfig != nil {
		return
	}
	globalConfig = new(Config)
	globalConfig.RaftCluster.Servers = make([]Server, 0)
	mu = new(sync.Mutex)

	permissions := os.FileMode(0755)
	if err := os.MkdirAll("data/persistency/", permissions); err != nil {
		panic(err)
	}
	globalConfig.Log = logger.MakeLogger()
}

func InitWithCommands(f *Flags, log logger.Logger) {
	mu.Lock()
	defer mu.Unlock()

	globalConfig.ID = f.ID
	globalConfig.Port = f.Port
	globalConfig.RaftCluster.ElectionTimeout = f.ElectionTimeout
	globalConfig.RaftCluster.HeartbeatInterval = f.HeartbeatInterval
	globalConfig.RaftCluster.EnableK8s = f.EnableK8s
	globalConfig.RaftCluster.ServerCount = f.ServerCount

	for id, addr := range f.Servers {
		host, port, _ := net.SplitHostPort(addr)
		globalConfig.RaftCluster.Servers = append(globalConfig.RaftCluster.Servers,
			Server{
				ID:   id,
				Host: host,
				Port: port,
			},
		)
	}

	globalConfig.Log = log
}

func InitWithYaml(log logger.Logger) {
	mu.Lock()
	defer mu.Unlock()
	file, err := os.ReadFile("raft.yaml")
	if err != nil {
		globalConfig.Log.With("err", err).Warn("Failed to open YAML file", "err", err)
		return
	}

	globalConfig.Log = log
	if err = yaml.Unmarshal(file, &globalConfig); err != nil {
		globalConfig.Log.With("err", err).Info("Failed to parse YAML file", "err")
		return
	}
}

func Get() Config {
	mu.Lock()
	defer mu.Unlock()
	cfg := *globalConfig
	copy(cfg.RaftCluster.Servers, globalConfig.RaftCluster.Servers)

	return cfg
}

func Set(cfg Config) {
	mu.Lock()
	defer mu.Unlock()
	var tmp []Server = make([]Server, 0)

	copy(tmp, cfg.RaftCluster.Servers)
	cfg.RaftCluster.Servers = tmp

	*globalConfig = cfg
}

func RemoveNode(key string) {
	cfg := Get()

	for i, srv := range cfg.RaftCluster.Servers {
		if srv.ID == key {
			cfg.RaftCluster.Servers = append(cfg.RaftCluster.Servers[:i], cfg.RaftCluster.Servers[i+1:]...)
			cfg.RaftCluster.ServerCount -= 1
			break
		}
	}

	Set(cfg)
}

func AddNode(key string, addr string) {
	cfg := Get()
	host, port, _ := net.SplitHostPort(addr)
	cfg.RaftCluster.Servers = append(cfg.RaftCluster.Servers, Server{
		ID:   key,
		Host: host,
		Port: port,
	})
	cfg.RaftCluster.ServerCount += 1

	Set(cfg)
}
