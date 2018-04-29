package lancher

import (
	"fmt"

	"github.com/breakD/algorithms/dstruct/stack"
)

// SimpleFixedStack example
func LanchLLStack() {
	llStack := &stack.LinkListStack{}
	count := 10
	for index := 0; index < count; index++ {
		llStack.Push(index)
	}
	llStack.Traverse()

	for index := 0; index < count; index++ {
		llStack.Pop()
		fmt.Printf("Pop result: stack size: %d, isEmpty: %t\n", llStack.Size(), llStack.IsEmpty())
	}
}
