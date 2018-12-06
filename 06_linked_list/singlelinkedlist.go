package linkedlist

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

type SingleLinkedList struct {
	head   *Node
	length int
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{}
}

func (s *SingleLinkedList) FindByValue(value interface{}) *Node {
	p := s.head
	for p != nil && p.data != value {
		p = p.next
	}
	return p
}

func (s *SingleLinkedList) FindByIndex(index int) *Node {
	p := s.head
	position := 0
	for p != nil && position != index {
		p = p.next
		position++
	}
	return p
}

func (s *SingleLinkedList) InsertValueToTail(value interface{}) *Node {
	node := &Node{data: value}
	s.InsertNodeToTail(node)
	return node
}

func (s *SingleLinkedList) InsertNodeToTail(node *Node) {
	if err := checkNode(node); err != nil {
		return
	}

	p := s.head
	for p.next != nil {
		p = p.next
	}
	p.next = node
	s.length++
}

func (s *SingleLinkedList) InsertValueToHead(value interface{}) *Node {
	node := &Node{data: value}
	s.InsertNodeToHead(node)
	return node
}

func (s *SingleLinkedList) InsertNodeToHead(node *Node) {
	if err := checkNode(node); err != nil {
		return
	}

	node.next = s.head
	s.head = node
	s.length++
}
func (s *SingleLinkedList) InsertValueAfter(targetNode *Node, value interface{}) *Node {
	newNode := &Node{data: value}
	s.InsertNodeAfter(targetNode, newNode)
	return newNode
}

func (s *SingleLinkedList) InsertNodeAfter(targetNode *Node, newNode *Node) {
	if err := checkNode(targetNode, newNode); err != nil {
		return
	}

	newNode.next = targetNode.next
	targetNode.next = newNode
	s.length++
}

func (s *SingleLinkedList) InsertValueBefore(targetNode *Node, value interface{}) *Node {
	newNode := &Node{data: value}
	s.InsertNodeBefore(targetNode, newNode)
	return newNode
}

func (s *SingleLinkedList) InsertNodeBefore(targetNode *Node, newNode *Node) {
	if err := checkNode(s.head, targetNode, newNode); err != nil {
		return
	}

	// 如果当前节点就是头节点，直接将新节点插入头节点
	if s.head == targetNode {
		s.InsertNodeToHead(newNode)
		return
	}

	p := s.head
	for p.next != nil && p.next != targetNode {
		p = p.next
	}

	// 如果 p == nil 说明目标节点不存在
	if err := checkNode(p); err != nil {
		return
	}

	newNode.next = targetNode
	p.next = newNode
	s.length++
}

func (s *SingleLinkedList) DeleteByNode(node *Node) {
	if err := checkNode(s.head, node); err != nil {
		return
	}

	// 如果删除的是头节点
	if s.head == node {
		s.head = s.head.next
		s.length--
		return
	}

	current := s.head
	for current.next != nil && current.next != node {
		current = current.next
	}

	if err := checkNode(current.next); err != nil {
		return
	}

	current.next = node.next
	s.length--
}

func (s *SingleLinkedList) String() string {
	var listStr string
	current := s.head
	for current != nil {
		listStr += fmt.Sprintf("%d\n", current.data)
		current = current.next
	}
	return listStr
}

func (s *SingleLinkedList) Traverse(handleFunc func(a, b interface{}) error, wants ...interface{}) error {
	p := s.head
	for _, want := range wants {
		if p == nil {
			return fmt.Errorf("node is nil error")
		}
		if err := handleFunc(p.data, want); err != nil {
			return err
		}
		p = p.next
	}
	return nil
}

func checkNode(nodes ...*Node) error {
	for _, node := range nodes {
		if node == nil {
			err := fmt.Errorf("node is nil error")
			fmt.Println(err)
			return err
		}
	}
	return nil
}
