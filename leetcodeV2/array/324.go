package array

import "sort"

/*
	solution 1: sort and seperate the array to large and small
		get 1 in turn
	solution 2: (1 + 2*index) % (n | 1)?
*/

func wiggleSort(nums []int) {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(tmp)
	m := (len(nums) + 1) >> 1

	for i, j := m-1, 0; i >= 0; i, j = i-1, j+2 {
		nums[j] = tmp[i]
	}
	for i, j := len(tmp)-1, 1; i >= m; i, j = i-1, j+2 {
		nums[j] = tmp[i]
	}
}
