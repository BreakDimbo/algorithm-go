package dynamicp

import "github.com/breakD/algorithms/leetcodeV2/util"

/*
	solution 1: dp[][], 0 1 for max and min
*/

func maxProduct(nums []int) int {
	const max, min = 0, 1
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][min] = 0
	dp[0][max] = nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {

		dp[i][min] = util.Min(dp[i-1][min]*nums[i], dp[i-1][max]*nums[i], nums[i])
		dp[i][max] = util.Max(dp[i-1][max]*nums[i], dp[i-1][max]*nums[i], nums[i])

		res = util.Max(dp[i][max], res)
	}

	return res
}
