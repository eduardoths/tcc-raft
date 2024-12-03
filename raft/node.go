package raft

type Node struct {
	Address string
}

func NewNode(address string) *Node {
	node := new(Node)
	node.Address = address
	return node
}
