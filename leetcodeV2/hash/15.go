package hash

import "sort"

/*
	solution 1: brute force O(n^3)
	solution 2: 两层循环 + map + 判重不好实现
	solution 2: 先排序，开始一层循环k，两个游标 i,j 从两端移动, 判断 num[k]+num[i]+num[j] > 0
*/

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := range nums {
		// jump repeated elements
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// jump repeated elements
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}

	return res
}
