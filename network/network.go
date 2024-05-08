package network

import (
	"math/rand"
	"time"

	"github.com/devlongs/gossiper/types"
)

type Network struct {
    Nodes []*types.Node
}

func GossipInterval() time.Duration {
    return time.Duration(rand.Intn(1000)) * time.Millisecond
}

func SelectRandomNeighbor(n *types.Node) *types.Node {
    neighborIndex := rand.Intn(len(n.Neighbors))
    return n.Neighbors[neighborIndex]
}