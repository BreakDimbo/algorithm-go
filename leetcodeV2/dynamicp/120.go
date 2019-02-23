package dynamicp

import "github.com/breakD/algorithms/leetcodeV2/util"

/*
	solution 1: dp
		dp[i][j] = min(dp[i-1][j-1],dp[i-1][j+1])
	solution 2: recursive
*/

/*
func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}

	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(triangle)-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = util.Min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}
	return util.Min(dp[len(triangle)-1]...)
}
*/

func minimumTotal(triangle [][]int) int {
	mini := triangle[len(triangle)-1]

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			mini[j] = util.Min(mini[j], mini[j+1]) + triangle[i][j]
		}
	}
	return mini[0]
}
