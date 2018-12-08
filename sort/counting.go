package sort

func countingSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}

	// find the max value
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}

	// init the counting bucket[0...max]
	bucket := make([]int, max+1)

	// count
	for _, v := range a {
		bucket[v]++
	}

	// accumulate
	// it's not good to modify the array when iteratin it
	for i, v := range bucket {
		if i == 0 {
			continue
		}

		bucket[i] = bucket[i-1] + v
	}

	// init the array for the sorted elements
	sorted := make([]int, n)

	// put the elements into the sorted
	// crucial part
	for i := n - 1; i >= 0; i-- {
		index := bucket[a[i]] - 1
		sorted[index] = a[i]
		bucket[a[i]]--
	}

	copy(a, sorted)
}
