package types

import "sync"

type Node struct {
    ID        int
    Neighbors []*Node
    State     State
    Mutex     sync.Mutex
}