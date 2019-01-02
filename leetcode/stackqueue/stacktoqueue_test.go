package stackqueue

import (
	"fmt"
	"testing"
)

func TestIntStack_Push(t *testing.T) {
	type args struct {
		e int
	}
	tests := []struct {
		name string
		s    *IntStack
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.e)
		})
	}
}

func TestMyQueue_Push(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Printf("%v\n", queue.Peek())
	fmt.Printf("%v\n", queue.Pop())
	// queue.Pop()
	// fmt.Printf("%v", queue)
	// queue.Empty()
	// fmt.Printf("%v", queue)
}
