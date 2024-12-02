package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/eduardoths/tcc-raft/raft"
	"github.com/eduardoths/tcc-raft/server"
	yaml "gopkg.in/yaml.v3"
)

func init() {
	permissions := os.FileMode(0755)
	if err := os.MkdirAll("data/persistency/", permissions); err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.ReadFile("raft.yaml")
	if err != nil {
		log.Fatalf("Failed to open YAML file: %v", err)
	}

	var config Config
	if err = yaml.Unmarshal(file, &config); err != nil {
		log.Fatalf("Failed to parse YAML file: %v", err)
	}

	nodes := make(map[raft.ID]*raft.Node)
	for i, s := range config.RaftCluster.Servers {
		if !strings.HasPrefix(s.Port, ":") {
			config.RaftCluster.Servers[i].Port = fmt.Sprintf(":%s", s.Port)
			s.Port = config.RaftCluster.Servers[i].Port
		}
		nodes[s.ID] = raft.NewNode(fmt.Sprintf("[::]%s", s.Port))
	}

	b, _ := json.Marshal(config)
	fmt.Printf("%s", string(b))

	var wg sync.WaitGroup
	for _, s := range config.RaftCluster.Servers {
		wg.Add(1)
		go func(srv Server) {
			defer wg.Done()
			server := server.CreateServer(
				raft.MakeRaft(srv.ID, nodes),
			)
			server.Start(srv.Port)
		}(s)
	}

	wg.Wait()
}
