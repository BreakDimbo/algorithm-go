package dynamicp

import "github.com/breakD/algorithms/leetcodeV2/util"

/*
	solution 1: dp
	solution 2: dfs
*/

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i][0] = util.Max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return util.Max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}
