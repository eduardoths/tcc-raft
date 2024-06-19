package main

import (
	"sync"

	"github.com/eduardoths/tcc-raft/raft"
)

func main() {
	nodes := make(map[raft.ID]*raft.Node)
	nodes["A"] = raft.NewNode("[::]:8080")
	nodes["B"] = raft.NewNode("[::]:8081")
	nodes["C"] = raft.NewNode("[::]:8082")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		server := raft.CreateServer(
			raft.MakeRaft("A", nodes),
		)
		server.Start(":8080")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		server := raft.CreateServer(
			raft.MakeRaft("B", nodes),
		)
		server.Start(":8081")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		server := raft.CreateServer(
			raft.MakeRaft("C", nodes),
		)
		server.Start(":8082")
	}()

	wg.Wait()
}
