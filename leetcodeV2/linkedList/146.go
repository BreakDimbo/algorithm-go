package linkedList

import (
	"container/list"
)

/*
	solution 1: linkedList
	solution 2: use hash to accelerate
*/

/*
type ListNode struct {
	Next *ListNode
	Prev *ListNode
	Key  int
	Val  int
}

type LRUCache struct {
	head     *ListNode
	tail     *ListNode
	loc      map[int]*ListNode
	capacity int
	length   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		loc:      make(map[int]*ListNode),
	}
}

// traverse the linkedlist find the node, delete then add it to the head
func (this *LRUCache) Get(key int) int {
	if node, ok := this.loc[key]; ok {
		this.moveToHead(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.loc[key]; ok {
		node.Val = value
		this.moveToHead(node)
	} else {
		newNode := &ListNode{
			Next: this.head,
			Key:  key,
			Val:  value,
		}
		if this.head != nil {
			this.head.Prev = newNode
		}

		if this.tail == nil {
			this.tail = newNode
		}

		this.head = newNode

		this.loc[key] = newNode

		if this.length < this.capacity {
			this.length++
		} else {
			// find the last node and delete it
			oldLast := this.tail
			delete(this.loc, oldLast.Key)
			newLast := oldLast.Prev
			newLast.Next = nil
			this.tail = newLast
		}
	}
}

func (this *LRUCache) moveToHead(node *ListNode) {
	if node != this.head {
		if node == this.tail {
			this.tail = node.Prev
		}
		node.Prev.Next = node.Next
		if node.Next != nil {
			node.Next.Prev = node.Prev
		}
		node.Next = this.head
		node.Prev = nil
		this.head.Prev = node
		this.head = node
	}
}
*/

type LRUCache struct {
	head     *list.List
	loc      map[int]*list.Element
	capacity int
	length   int
}

type Node struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		head:     list.New(),
		loc:      make(map[int]*list.Element),
	}
}

// traverse the linkedlist find the node, delete then add it to the head
func (this *LRUCache) Get(key int) int {
	if node, ok := this.loc[key]; ok {
		this.head.MoveToFront(node)
		return node.Value.(*Node).Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.loc[key]; ok {
		// if node exist, update value and move node to head
		node.Value.(*Node).Value = value
		this.head.MoveToFront(node)
	} else {
		element := this.head.PushFront(&Node{key, value})
		this.loc[key] = element
		if this.length < this.capacity {
			this.length++
		} else {
			// find the last node and delete it
			last := this.head.Back()
			delete(this.loc, last.Value.(*Node).Key)
			this.head.Remove(last)
		}
	}
}
