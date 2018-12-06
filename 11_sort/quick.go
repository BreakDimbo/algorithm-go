package sort

func quickSort(a []int) {
	if len(a) <= 1 {
		return
	}
	qSort(a, 0, len(a)-1)
}

func qSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}

	p := partition(a, lo, hi)
	qSort(a, lo, p-1)
	qSort(a, p+1, hi)
}

func partition(a []int, lo, hi int) int {
	i, j := lo, hi+1
	pivot := a[lo]
	for {
		for {
			i++
			if a[i] >= pivot || i == hi {
				break
			}
		}

		for {
			j--
			if a[j] <= pivot || j == lo {
				break
			}
		}

		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[lo], a[j] = a[j], a[lo]
	return j
}
