package search

import (
	"reflect"
	"testing"
)

func TestNewGraph(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want *Graph
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGraph(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_Add(t *testing.T) {
	graph := NewGraph(10)
	graph.Add(1, 3)
	graph.Add(1, 5)
	graph.Print()
}

func TestGraph_BFS(t *testing.T) {
	graph := NewGraph(8)
	graph.Add(0, 1)
	graph.Add(0, 3)
	graph.Add(1, 2)
	graph.Add(1, 4)
	graph.Add(3, 4)
	graph.Add(2, 5)
	graph.Add(4, 5)
	graph.Add(4, 6)
	graph.Add(6, 7)
	graph.Add(5, 7)
	graph.BFS(0, 7)
}

func TestGraph_DFS(t *testing.T) {
	graph := NewGraph(9)
	graph.Add(1, 2)
	graph.Add(1, 4)
	graph.Add(2, 5)
	graph.Add(4, 5)

	graph.Add(2, 3)
	graph.Add(3, 6)
	graph.Add(5, 6)

	graph.Add(5, 7)
	graph.Add(7, 8)
	graph.Add(6, 8)
	graph.DFS(1, 7)
}
