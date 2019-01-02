package stackqueue

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(a, b int) bool { return h[a] < h[b] }
func (h IntHeap) Swap(a, b int)      { h[a], h[b] = h[b], h[a] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type KthLargest struct {
	heap   IntHeap
	length int
}

func PConstructor(k int, nums []int) KthLargest {
	h := (*IntHeap)(&nums)
	heap.Init(h)
	return KthLargest{
		heap:   *h,
		length: k,
	}
}

func (this *KthLargest) Add(val int) int {
	for len(this.heap) > this.length {
		heap.Pop(&this.heap)
	}

	if len(this.heap) < this.length {
		heap.Push(&this.heap, val)
		return this.heap[0]
	}

	if val > this.heap[0] {
		heap.Pop(&this.heap)
		heap.Push(&this.heap, val)
	}
	return this.heap[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
