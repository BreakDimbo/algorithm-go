package binary

/*
	solution 1: brute force / dfs
	solution 2: Backtracking swap
*/

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	var b func([]int, int)

	b = func(nums []int, first int) {
		if first == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			ans = append(ans, tmp)
			return
		}

		// 注意 i 从 first 开始
		for i := first; i < len(nums); i++ {
			nums[i], nums[first] = nums[first], nums[i]
			b(nums, first+1)
			nums[i], nums[first] = nums[first], nums[i]
		}
	}

	b(nums, 0)
	return ans
}
