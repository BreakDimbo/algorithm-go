package dynamic

import (
	"math"
)

func maxProfit(k int, prices []int) int {
	if k == 0 || prices == nil {
		return 0
	}

	/*
		init state: max profit
		mp[i][k][j]
		i: nth day
		k: number of transaction
		j: 0, 1 if hold a stock
	*/
	n := len(prices)

	if k > n>>1 {
		return quickSolve(prices)
	}

	mp := make([][][]int, n)
	if k == 1000000000 {
		k = n / 2
	}
	for i := range mp {
		mp[i] = make([][]int, k+1)
		for j := range mp[i] {
			mp[i][j] = make([]int, 2)
			if i == 0 {
				if j == 0 {
					mp[i][j][0] = 0
					mp[i][j][1] = -prices[0]
				} else {
					mp[i][j][0] = math.MinInt32
					mp[i][j][1] = math.MinInt32
				}
			}
		}
	}
	res := 0

	for i := 1; i < n; i++ {
		for j := 0; j <= k; j++ {
			if j == 0 {
				mp[i][j][0] = mp[i-1][j][0]
				mp[i][j][1] = max(mp[i-1][j][1], mp[i-1][j][0]-prices[i])
			} else {
				mp[i][j][0] = max(mp[i-1][j][0], mp[i-1][j-1][1]+prices[i])
				mp[i][j][1] = max(mp[i-1][j][1], mp[i-1][j][0]-prices[i])
			}
			res = max(mp[i][j][0], mp[i][j][1], res)
		}
	}
	return res
}

// for large test case
func quickSolve(prices []int) int {
	len, profit := len(prices), 0

	for i := 1; i < len; i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
