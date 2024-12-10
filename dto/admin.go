package dto

import (
	"fmt"

	pb "github.com/eduardoths/tcc-raft/proto"
)

type AddNodeArgs struct {
	ID   string `json:"id"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (ana AddNodeArgs) Addr() string {
	return fmt.Sprintf("%s:%d", ana.Host, ana.Port)
}

type SetNodesArgs struct {
	Nodes map[ID]string
}

func SetNodesArgsFromProto(proto *pb.SetNodesArgs) SetNodesArgs {
	nodes := proto.GetNodes()
	nodesArgs := SetNodesArgs{
		Nodes: make(map[ID]string, len(nodes)),
	}
	for k, v := range nodes {
		nodesArgs.Nodes[k] = v.GetAddress()
	}

	return nodesArgs
}
