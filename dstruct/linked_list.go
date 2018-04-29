package dstruct

import "github.com/breakD/algorithms/item"

type Node struct {
	Item item.Item
	Next *Node
}

func NewNode() *Node {
	return &Node{}
}
