package sort

/*
heap sort
大堆
小堆
自上而下
自下而上
*/

type Heap struct {
	data   []int
	length int
	count  int
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		data:   make([]int, capacity),
		length: capacity,
		count:  0,
	}
}

// down to top
func (h *Heap) Insert(value int) {
	if h.count >= h.length {
		return
	}
	h.count++
	h.data[h.count] = value

	i := h.count
	for i/2 > 0 && h.data[i] > h.data[i/2] {
		swap(h.data, i, i/2)
		i = i / 2
	}
}

func (h *Heap) Pop() {
	if h.count == 0 {
		return
	}
	h.data[1] = h.data[h.count]
	h.count--
	heapify(h.data, h.length, h.count)
}

// top to down
func heapify(a []int, length, i int) {
	for {
		maxPos := i
		if 2*i <= length && a[i] < a[2*i] {
			maxPos = 2 * i
		}
		if 2*i+1 <= length && a[maxPos] < a[2*i+1] {
			maxPos = 2*i + 1
		}
		if maxPos == i {
			break
		}
		swap(a, i, maxPos)
		i = maxPos
	}
}

func HeapSort(a []int) {
	k := len(a) - 1
	buildHeap(a, k)

	for k > 1 {
		swap(a, 1, k)
		k--
		heapify(a, k, 1)
	}
}

func buildHeap(a []int, n int) {
	for i := n / 2; i > 0; i-- {
		heapify(a, n, i)
	}
}

func swap(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}
