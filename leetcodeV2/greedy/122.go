package greedy

import "github.com/breakD/algorithms/leetcodeV2/util"

/*
	solution 1: greedy
		if prices[i-1] > prices[i] sell
		if prices[i] < prices[i+1] buy
	solution 2: DFS
		findProfit(prices, hold, profit)
	solution 3: DP
		DP : dp[i][j]: 第 i 天的利润，是否持有股票
*/

/*
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := 0
	for i := 0; i < len(prices); i++ {
		if i > 0 && prices[i] > prices[i-1] {
			profit += prices[i]
		}
		if i < len(prices)-1 && prices[i] < prices[i+1] {
			profit -= prices[i]
		}
	}
	return profit
}
*/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][1] = -prices[0]
	dp[0][0] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = util.Max(dp[i-1][1]+prices[i], dp[i-1][0])
		dp[i][1] = util.Max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][0]
}
