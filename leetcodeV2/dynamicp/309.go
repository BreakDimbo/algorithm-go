package dynamicp

import "github.com/breakD/algorithms/leetcodeV2/util"

/*
	solution 1: dp
*/

func maxProfitXXX(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	// 0: hold 1: sell 2: cooldown 3: unhold

	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0
	dp[0][3] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = util.Max(dp[i-1][2]-prices[i], dp[i-1][0], dp[i-1][3]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = dp[i-1][1]
		dp[i][3] = util.Max(dp[i-1][2], dp[i-1][3])
	}
	return util.Max(dp[len(prices)-1][0], dp[len(prices)-1][1], dp[len(prices)-1][2], dp[len(prices)-1][3])
}
