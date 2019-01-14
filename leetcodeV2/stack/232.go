package stack

type MyQueue struct {
	in  *Stack
	out *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		in:  New(),
		out: New(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.in.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.out.Len() == 0 {
		for v := this.in.Pop(); v != nil; v = this.in.Pop() {
			this.out.Push(v)
		}
	}
	return this.out.Pop().(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.out.Len() == 0 {
		for v := this.in.Pop(); v != nil; v = this.in.Pop() {
			this.out.Push(v)
		}
	}
	return this.out.Peek().(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.in.Len() == 0 && this.out.Len() == 0
}
