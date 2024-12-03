package raft

import (
	"context"

	"github.com/eduardoths/tcc-raft/dto"
)

func (r *Raft) SetNodes(ctx context.Context, args dto.SetNodesArgs) error {
	r.nodexMutex.Lock()
	defer r.nodexMutex.Unlock()

	m := make(map[ID]*Node, len(args.Nodes))
	for k, v := range args.Nodes {
		m[k] = &Node{Address: v}
	}

	r.setNodes(m)

	return nil
}

func (r *Raft) Shutdown(ctx context.Context) error {
	r.stop <- struct{}{}
	return nil
}
