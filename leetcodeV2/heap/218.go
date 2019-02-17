package heap

import (
	"container/heap"
	"sort"
)

/*
	solution 1: min Heap
*/

type Item struct {
	value int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].value > pq[j].value }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq PriorityQueue) Top() *Item {
	if len(pq) > 0 {
		return pq[0]
	}
	return nil
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	v := old[n-1]
	v.index = -1
	*pq = old[:n-1]
	return v
}

func getSkyline(buildings [][]int) [][]int {
	result := make([][]int, 0)
	height := make([][]int, 0)

	for _, b := range buildings {
		height = append(height, []int{b[0], -b[2]})
		height = append(height, []int{b[1], b[2]})
	}

	// handle ties by put start before end
	sort.Slice(height, func(i, j int) bool {
		if height[i][0] != height[j][0] {
			return height[i][0] < height[j][0]
		}
		return height[i][1] < height[j][1]
	})

	pq := &PriorityQueue{&Item{value: 0}}
	heap.Init(pq)

	prev := 0

	for _, h := range height {
		// if h is start
		if h[1] < 0 {
			heap.Push(pq, &Item{value: -h[1]})
		} else {
			// remove
			for i, v := range *pq {
				if h[1] == v.value {
					heap.Remove(pq, i)
					break
				}
			}
		}

		curMax := pq.Top().value

		if prev != curMax {
			result = append(result, []int{h[0], curMax})
			prev = curMax
		}
	}
	return result
}
