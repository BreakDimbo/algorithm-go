package stack

import (
	"fmt"

	"github.com/breakD/algorithms/dstruct"
	"github.com/breakD/algorithms/item"
)

type LinkListStack struct {
	firstNode *dstruct.Node
	n         int // number of item
}

func (l *LinkListStack) IsEmpty() bool {
	return l.firstNode == nil
}

func (l *LinkListStack) Size() int {
	return l.n
}

func (l *LinkListStack) Push(i item.Item) {
	oldFirst := l.firstNode
	n := dstruct.NewNode()
	n.Item = i
	n.Next = oldFirst
	l.firstNode = n
	l.n += 1
}

func (l *LinkListStack) Pop() (i item.Item) {
	i = l.firstNode.Item
	l.firstNode = l.firstNode.Next
	l.n -= 1
	return
}

func (l *LinkListStack) Traverse() {
	fmt.Printf("Hi we are going to start traverse the stack which size is: %d\n", l.Size())
	for x := l.firstNode; x != nil; x = x.Next {
		fmt.Printf("the item is: %v\n", x.Item)
	}
}
