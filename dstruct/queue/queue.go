package queue

import (
	"fmt"

	"github.com/breakD/algorithms/dstruct"
	"github.com/breakD/algorithms/item"
)

type Queue struct {
	firstNode *dstruct.Node
	lastNode  *dstruct.Node
	n         int
}

func (q *Queue) IsEmpty() bool {
	return q.firstNode == nil
}

func (q *Queue) Size() int {
	return q.n
}

func (q *Queue) Enqueue(i item.Item) {
	oldLast := q.lastNode
	node := dstruct.NewNode()
	node.Item = i
	node.Next = nil
	q.lastNode = node
	if q.IsEmpty() {
		q.firstNode = node
	} else {
		oldLast.Next = node
	}

	q.n += 1
}

func (q *Queue) Dequeue() (i item.Item) {
	i = q.firstNode.Item
	q.firstNode = q.firstNode.Next
	if q.IsEmpty() {
		q.lastNode = nil
	}
	q.n -= 1
	return
}

func (q *Queue) Traverse() {
	fmt.Printf("start traversing the queue which size is: %d\n", q.n)
	for x := q.firstNode; x != nil; x = x.Next {
		fmt.Printf("the item is: %v\n", x.Item)
	}
}
