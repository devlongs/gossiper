package node

import (
	"time"

	"github.com/devlongs/gossiper/network"
	"github.com/devlongs/gossiper/types"
)

func Gossip(n *types.Node) {
    for {
        // Sleep for a random interval
        time.Sleep(network.GossipInterval())

        // Select a random neighbor
        neighbor := network.SelectRandomNeighbor(n)

        // Lock the current node's state
        n.Mutex.Lock()
        currentState := n.State
        n.Mutex.Unlock()

        // Exchange state with the selected neighbor
        neighbor.Mutex.Lock()
        if currentState.Value > neighbor.State.Value {
            neighbor.State = currentState
        }
        neighbor.Mutex.Unlock()
    }
}