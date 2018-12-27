package search

import (
	"reflect"
	"testing"
)

func TestNewTree(t *testing.T) {
	tests := []struct {
		name string
		want *Tree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNode(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Add(t *testing.T) {
	tree := NewTree()
	data := []int{2, 3, 1, 9, 6, 4, 21, 17, 12, 32}
	for _, v := range data {
		tree.Add(v)
	}
	tree.PrintPreOrder()
	tree.PrintInOrder()
	tree.PrintPostOrder()
}

func TestTree_Del(t *testing.T) {
	tree := NewTree()
	data := []int{2, 3, 1, 9, 6, 4, 21, 17, 12, 32}
	for _, v := range data {
		tree.Add(v)
	}

	tree.Del(2)
	tree.PrintInOrder()

	tree.Del(11)

	tree.Del(21)
	tree.Del(17)
	tree.PrintInOrder()
}

func TestTree_Find(t *testing.T) {
	tree := NewTree()
	data := []int{2, 3, 1, 9, 6, 4, 21, 17, 12, 32}
	for _, v := range data {
		tree.Add(v)
	}

	t.Log(tree.Find(32).data)
}
