package dynamicp

import (
	"math"

	"github.com/breakD/algorithms/leetcodeV2/util"
)

/*
	solution 1: dp[i][j][k]
		i : days
		j : number of transactions
		k : number of hold stock
*/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][][]int, len(prices))
	for i := range dp {
		dp[i] = make([][]int, 3)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
			if i == 0 && j != 0 {
				dp[i][j][0] = math.MinInt32
				dp[i][j][1] = math.MinInt32
			}
		}
	}
	dp[0][0][0] = 0
	dp[0][0][1] = -prices[0]
	max := 0

	for i := 1; i < len(prices); i++ {
		for j := 0; j < 3; j++ {
			if j == 0 {
				dp[i][j][0] = dp[i-1][j][0]
			} else {
				dp[i][j][0] = util.Max(dp[i-1][j][0], dp[i-1][j-1][1]+prices[i])
			}
			dp[i][j][1] = util.Max(dp[i-1][j][1], dp[i-1][j][0]-prices[i])
			max = util.Max(max, dp[i][j][0], dp[i][j][1])
		}
	}
	return max
}
