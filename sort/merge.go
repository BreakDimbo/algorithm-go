package sort

func mergeSort(a []int) {
	mSort(a, 0, len(a)-1)
}

func mSort(a []int, lo, hi int) {
	// 递归终止
	if lo >= hi {
		return
	}

	mid := lo + (hi-lo)/2

	mSort(a, lo, mid)
	mSort(a, mid+1, hi)

	merge(a, lo, mid, hi)
}

func merge(a []int, lo, mid, hi int) {
	i, j := lo, mid+1
	tmp := make([]int, len(a))
	for k := lo; k <= hi; k++ {
		tmp[k] = a[k]
	}

	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = tmp[j]
			j++
		} else if j > hi {
			a[k] = tmp[i]
			i++
		} else if tmp[i] < tmp[j] {
			a[k] = tmp[i]
			i++
		} else {
			a[k] = tmp[j]
			j++
		}
	}
}
