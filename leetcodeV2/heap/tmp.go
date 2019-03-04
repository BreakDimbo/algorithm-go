package heap

import "container/heap"

type IHeap []int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IHeap) Top() int           { return h[0] } // if len(h.data) == 0 it will panic

func (h *IHeap) Pop() interface{} {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

type KthLargest struct {
	h   *IHeap
	cap int
}

func Constructor(k int, nums []int) KthLargest {
	h := IHeap(nums)
	heap.Init(&h)
	return KthLargest{
		h:   &h,
		cap: k,
	}
}

func (this *KthLargest) Add(val int) int {
	for this.h.Len() > this.cap {
		heap.Pop(this.h)
	}

	if this.h.Len() == this.cap && val > this.h.Top() {
		heap.Pop(this.h)
	}

	if this.h.Len() < this.cap {
		heap.Push(this.h, val)
	}
	return this.h.Top()
}
