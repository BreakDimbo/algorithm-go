package bdfs

import "sort"

/*
	solution 1: brute force
*/

func subsets(nums []int) [][]int {
	sort.Ints(nums)

	return dfs([]int{}, nums, 0)
}

func dfs(curAns, nums []int, start int) [][]int {
	ans := [][]int{}
	ans = append(ans, curAns)
	for i := start; i < len(nums); i++ {
		nextAns := []int{}
		nextAns = append(nextAns, curAns...)
		nextAns = append(nextAns, nums[i])
		ans = append(ans, dfs(nextAns, nums, i+1)...)
	}
	return ans
}
