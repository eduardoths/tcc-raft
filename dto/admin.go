package dto

import (
	pb "github.com/eduardoths/tcc-raft/proto"
)

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
