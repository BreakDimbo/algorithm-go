package others

import (
	"container/heap"
)

/*
	solution 1: use a sorted array O(nlg)
	solution 2: use 2 heap - minHeap and maxHeap for half data seperately
	solution 3: BST, median is the root or one of its child
*/

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Top() int           { return h[0] }

func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Top() int           { return h[0] }

func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

type MedianFinder struct {
	lo *MaxHeap
	hi *MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	lo := &MaxHeap{}
	hi := &MinHeap{}

	heap.Init(lo)
	heap.Init(hi)
	return MedianFinder{
		lo: lo,
		hi: hi,
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.lo, num)

	heap.Push(this.hi, heap.Pop(this.lo))

	if this.lo.Len() < this.hi.Len() {
		heap.Push(this.lo, heap.Pop(this.hi))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lo.Len() > this.hi.Len() {
		return float64(this.lo.Top())
	}
	return float64(this.lo.Top()+this.hi.Top()) / 2
}
