package stackqueue

type IntStack []int

const Empty = -2147483648

func NewIntStack() *IntStack {
	var stack []int
	return (*IntStack)(&stack)
}

func (s *IntStack) Push(e int) {
	*s = append(*s, e)
}

func (s *IntStack) Pop() int {
	if len(*s) == 0 {
		return Empty
	}

	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *IntStack) Peek() int {
	if len(*s) == 0 {
		return Empty
	}

	v := (*s)[len(*s)-1]
	return v
}

type MyQueue struct {
	intput *IntStack
	output *IntStack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		intput: NewIntStack(),
		output: NewIntStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.intput.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(*this.output) == 0 {
		for len(*this.intput) != 0 {
			e := this.intput.Pop()
			this.output.Push(e)
		}
	}
	return this.output.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(*this.output) == 0 {
		for len(*this.intput) != 0 {
			e := this.intput.Pop()
			this.output.Push(e)
		}
	}
	return this.output.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(*this.intput) == 0 && len(*this.output) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
