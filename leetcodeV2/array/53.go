package array

import "github.com/breakD/algorithms/leetcodeV2/util"

/*
	solution 1: brute force
	solution 2: dp
	solution 3: divide & conquer
*/

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		dp[i] = util.Max(nums[i], dp[i-1]+nums[i])
	}
	return util.Max(dp...)
}
