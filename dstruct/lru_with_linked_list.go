package dstruct

import (
	"github.com/breakD/algorithms/item"
)

type LinkedListLRU struct {
	head     *Node
	size     int
	capacity int
}

func NewLinkedListLRU(cap int) *LinkedListLRU {
	return &LinkedListLRU{
		capacity: cap,
	}
}

func (l *LinkedListLRU) Get(item item.Item) item.Item {
	// 如果没有头节点，将数据直接缓存至头节点
	if l.head == nil {
		l.head = NewNode()
		l.head.Item = item
		l.size += 1
		return item
	}

	var lastPreNode *Node
	for node := l.head; node.Next != nil; node = node.Next {
		// 如果缓存命中，将命中节点移动至头节点
		if node.Next.Item.Equal(item) {
			currentNode := node.Next
			node.Next = currentNode.Next
			currentNode.Next = l.head
			l.head = currentNode
			return item
		}
		lastPreNode = node
	}

	// 如果缓存未命中，判断当前缓存容量是否超过阈值
	// 如果超过阈值，删除尾部节点
	if l.size >= l.capacity {
		lastPreNode.Next = nil
	}

	// 将新节点加入头节点
	newNode := &Node{
		Item: item,
		Next: l.head,
	}
	l.head = newNode
	return item
}
