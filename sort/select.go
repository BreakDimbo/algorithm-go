package sort

func selectSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}
