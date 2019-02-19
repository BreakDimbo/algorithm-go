package others

/*
	solution 1: hash + heap
	solution 2: bucket sort
*/

/*
type Item struct {
	count   int
	element int
}

type ItemHeap []*Item

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h ItemHeap) Top() int           { return h[0].count }

func (h *ItemHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	h := &ItemHeap{}
	heap.Init(h)

	for element, count := range m {
		if h.Len() < k {
			heap.Push(h, &Item{count, element})
		} else if count > h.Top() {
			heap.Pop(h)
			heap.Push(h, &Item{count, element})
		}
	}

	res := make([]int, h.Len())
	for i := range res {
		res[len(res)-i-1] = heap.Pop(h).(*Item).element
	}
	return res
}
*/

func topKFrequent(nums []int, k int) []int {
	bucket := make([][]int, len(nums)+1)
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	for ele, count := range m {
		bucket[count] = append(bucket[count], ele)
	}

	res := make([]int, 0)
	for i := len(bucket) - 1; i >= 0 && len(res) <= k; i-- {
		if len(bucket[i]) != 0 {
			res = append(res, bucket[i]...)
		}
	}
	return res
}
