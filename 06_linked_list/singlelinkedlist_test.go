package linkedlist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSingleLinkedList(t *testing.T) {
	list := NewSingleLinkedList()
	firstNode := list.InsertValueToHead(1)
	secondNode := list.InsertValueBefore(firstNode, 2)
	list.InsertValueAfter(secondNode, 3)
	list.InsertValueToHead(4)
	list.DeleteByNode(firstNode)

	err := list.Traverse(func(a, b interface{}) error {
		if a != b {
			return fmt.Errorf("a: %d is not equal b: %d", a, b)
		}
		return nil
	}, 4, 2, 3)

	t.Log(list)

	if err != nil {
		t.Error(err)
	}

	node := list.FindByValue(2)
	if node != secondNode {
		t.Errorf("want %d but get %d", secondNode.data, node.data)
	}
}

func TestNewSingleLinkedList(t *testing.T) {
	tests := []struct {
		name string
		want *SingleLinkedList
	}{
		// TODO: Add test cases.
		{"newlist", NewSingleLinkedList()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSingleLinkedList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSingleLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleLinkedList_FindByValue(t *testing.T) {
	type fields struct {
		head *Node
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head: tt.fields.head,
			}
			if got := s.FindByValue(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SingleLinkedList.FindByValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleLinkedList_FindByIndex(t *testing.T) {
	type fields struct {
		head *Node
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head: tt.fields.head,
			}
			if got := s.FindByIndex(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SingleLinkedList.FindByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleLinkedList_InsertValueToHead(t *testing.T) {
	type fields struct {
		head *Node
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head: tt.fields.head,
			}
			s.InsertValueToHead(tt.args.value)
		})
	}
}

func TestSingleLinkedList_InsertNodeToHead(t *testing.T) {
	type fields struct {
		head *Node
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head: tt.fields.head,
			}
			s.InsertNodeToHead(tt.args.node)
		})
	}
}

func TestSingleLinkedList_InsertNodeAfter(t *testing.T) {
	type fields struct {
		head *Node
	}
	type args struct {
		targetNode *Node
		newNode    *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head: tt.fields.head,
			}
			s.InsertNodeAfter(tt.args.targetNode, tt.args.newNode)
		})
	}
}

func TestSingleLinkedList_InsertValueAfter(t *testing.T) {
	type fields struct {
		head *Node
	}
	type args struct {
		targetNode *Node
		value      int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head: tt.fields.head,
			}
			s.InsertValueAfter(tt.args.targetNode, tt.args.value)
		})
	}
}

func TestSingleLinkedList_InsertNodeBefore(t *testing.T) {
	type fields struct {
		head *Node
	}
	type args struct {
		targetNode *Node
		newNode    *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head: tt.fields.head,
			}
			s.InsertNodeBefore(tt.args.targetNode, tt.args.newNode)
		})
	}
}

func TestSingleLinkedList_InsertValueBefore(t *testing.T) {
	type fields struct {
		head *Node
	}
	type args struct {
		targetNode *Node
		value      int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head: tt.fields.head,
			}
			s.InsertValueBefore(tt.args.targetNode, tt.args.value)
		})
	}
}

func TestSingleLinkedList_DeleteByNode(t *testing.T) {
	type fields struct {
		head *Node
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head: tt.fields.head,
			}
			s.DeleteByNode(tt.args.node)
		})
	}
}

func TestSingleLinkedList_String(t *testing.T) {
	type fields struct {
		head *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head: tt.fields.head,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("SingleLinkedList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkNode(t *testing.T) {
	type args struct {
		nodes []*Node
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkNode(tt.args.nodes...); (err != nil) != tt.wantErr {
				t.Errorf("checkNode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
