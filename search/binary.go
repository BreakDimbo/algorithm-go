package search

func bSearch(a []int, target int) int {
	n := len(a)
	lo := 0
	hi := n - 1

	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if a[mid] == target {
			return mid
		}

		if target < a[mid] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}

func bSearchR(a []int, target int) int {
	return bSearchInternally(a, 0, len(a)-1, target)
}

func bSearchInternally(a []int, lo, hi, target int) int {
	if lo > hi {
		return -1
	}

	mid := lo + ((hi - lo) >> 1)
	if target == a[mid] {
		return mid
	}

	if target < a[mid] {
		return bSearchInternally(a, lo, mid-1, target)
	} else {
		return bSearchInternally(a, mid+1, hi, target)
	}
}

func bSearchFirstEqual(a []int, target int) int {
	lo := 0
	hi := len(a) - 1

	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if target > a[mid] {
			lo = mid + 1
		} else if target < a[mid] {
			hi = mid - 1
		} else {
			if mid == 0 || a[mid-1] != target {
				return mid
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

func bSearchLastEqual(a []int, target int) int {
	lo := 0
	hi := len(a) - 1

	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)

		if target > a[mid] {
			lo = mid + 1
		} else if target < a[mid] {
			hi = mid - 1
		} else {
			if mid == len(a)-1 || a[mid+1] != target {
				return mid
			} else {
				lo = mid + 1
			}
		}
	}
	return -1
}

func bSearchFirstGreaterOrEqual(a []int, target int) int {
	lo := 0
	hi := len(a) - 1

	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)

		if target <= a[mid] {
			if mid == 0 || a[mid-1] < target {
				return mid
			} else {
				hi = mid - 1
			}
		} else {
			lo = mid + 1
		}
	}
	return -1
}

func bSearchLastLessOrEqual(a []int, target int) int {
	lo := 0
	hi := len(a) - 1

	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if target < a[mid] {
			hi = mid - 1
		} else {
			if mid == len(a)-1 || a[mid+1] > target {
				return mid
			} else {
				lo = mid + 1
			}
		}
	}
	return -1
}
