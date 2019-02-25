package dynamicp

import "github.com/breakD/algorithms/leetcodeV2/util"

/*
	solution 1 : dp
		dp[i][j]: stands for word a ith rune and word b jth rune state
			when s[i] == s[j],  min steps
			s[i] add 1 or s[j] delete 1
			s[i] delete 1 or s[j] add 1
			s[i-1] replace or s[j-1] replace
*/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = util.Min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
