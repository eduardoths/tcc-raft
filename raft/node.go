package raft

type Node struct {
	Connect bool
	Address string
}

func NewNode(address string) *Node {
	node := new(Node)
	node.Address = address
	return node
}
