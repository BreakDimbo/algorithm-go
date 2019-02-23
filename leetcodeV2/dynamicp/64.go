package dynamicp

import "github.com/breakD/algorithms/leetcodeV2/util"

/*
	solution 1: dp
	solution 2: dfs/backtracing
*/

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[i]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][0] = dp[i-1][0] + grid[i][0]
			} else {
				dp[i][j] = util.Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}
