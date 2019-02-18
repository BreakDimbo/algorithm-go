package array

/*
	solution 1: brute force
	solution 2: merge sort the index
*/

/*
func countSmaller(nums []int) []int {
	res := make([]int, len(nums))
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				res[i]++
			}
		}
	}
	return res
}
*/

func countSmaller(nums []int) []int {
	count := make([]int, len(nums))
	index := make([]int, len(nums))
	for i := range index {
		index[i] = i
	}
	mergeSort(nums, 0, len(index)-1, index, count)
	return count
}

func mergeSort(nums []int, lo, hi int, index, count []int) {
	if lo >= hi {
		return
	}
	mid := lo + ((hi - lo) >> 1)
	mergeSort(nums, lo, mid, index, count)
	mergeSort(nums, mid+1, hi, index, count)

	mergeX(nums, lo, mid, mid+1, hi, index, count)
}

func mergeX(nums []int, l1, h1, l2, h2 int, index, count []int) {
	tmp := make([]int, h2-l1+1)
	i := 0
	rightCount := 0
	start := l1

	for l1 <= h1 || l2 <= h2 {
		if l1 > h1 {
			tmp[i] = index[l2]
			l2++
		} else if l2 > h2 {
			tmp[i] = index[l1]
			count[index[l1]] += rightCount
			l1++
		} else if nums[index[l1]] > nums[index[l2]] {
			tmp[i] = index[l2]
			rightCount++
			l2++
		} else {
			tmp[i] = index[l1]
			count[index[l1]] += rightCount
			l1++
		}
		i++
	}
	for i := range tmp {
		index[i+start] = tmp[i]
	}
}
