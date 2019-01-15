package priorityQueue

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Peek() interface{} {
	// 这里也可以直接使用 return (*h)[0]
	res := heap.Pop(h)
	heap.Push(h, res)
	return res
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type KthLargest struct {
	heap *IntHeap
	cap  int
}

func Constructor(k int, nums []int) KthLargest {
	h := IntHeap(nums)
	heap.Init(&h)
	return KthLargest{
		cap:  k,
		heap: &h,
	}
}

func (this *KthLargest) Add(val int) int {
	// 注意清空多余元素
	for this.heap.Len() > this.cap {
		heap.Pop(this.heap)
	}

	if this.heap.Len() < this.cap {
		heap.Push(this.heap, val)
	} else if val > this.heap.Peek().(int) {
		heap.Pop(this.heap)
		heap.Push(this.heap, val)
	}

	return this.heap.Peek().(int)
}
