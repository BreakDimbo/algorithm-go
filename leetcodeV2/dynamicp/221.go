package dynamicp

import "github.com/breakD/algorithms/leetcodeV2/util"

/*
	solution 1: dp
		Square ending at i,j is intersection of
		Square ending at i-1,j-1
		Square ending at i-1,j
		Square ending at i,j-1
*/

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxLen := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == byte('1') {
				dp[i][j] = util.Min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
				maxLen = util.Max(maxLen, dp[i][j])
			}
		}
	}
	return maxLen * maxLen
}
