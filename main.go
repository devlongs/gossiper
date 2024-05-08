package main

import (
	"fmt"
	"time"

	"github.com/devlongs/gossiper/node"
	"github.com/devlongs/gossiper/types"
)

func main() {
    // Create nodes
    nodes := make([]*types.Node, 5)
    for i := 0; i < 5; i++ {
        nodes[i] = &types.Node{ID: i, State: types.State{Value: i}}
    }

    // Set up neighbor connections
    nodes[0].Neighbors = []*types.Node{nodes[1], nodes[4]}
    nodes[1].Neighbors = []*types.Node{nodes[0], nodes[2], nodes[4]}
    nodes[2].Neighbors = []*types.Node{nodes[1], nodes[3]}
    nodes[3].Neighbors = []*types.Node{nodes[2], nodes[4]}
    nodes[4].Neighbors = []*types.Node{nodes[0], nodes[1], nodes[3]}

    // Start gossiping
    for _, n := range nodes {
        go node.Gossip(n)
    }

    // Wait and observe the state changes
    for {
        time.Sleep(2 * time.Second)
        fmt.Println("Node states:")
        for _, n := range nodes {
            fmt.Printf("Node %d: %d\n", n.ID, n.State.Value)
        }
        fmt.Println()
    }
}