package dynamicp

import "math"

/*
	solution 1: one pass
	solution 2: dp
		dp[i][0]
		dp[i][1]
*/

func maxProfitXX(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0

	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		} else if p-minPrice > maxProfit {
			maxProfit = p - minPrice
		}
	}
	return maxProfit
}
