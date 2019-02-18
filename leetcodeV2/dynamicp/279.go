package dynamicp

import (
	"math"

	"github.com/breakD/algorithms/leetcodeV2/util"
)

/*
	solution 1: dp
	solution 2: BFS
	solution 3: Mathematical
*/

func numSquares(n int) int {
	if n < 0 {
		return 0
	}
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		// no need to cal dp if i is perfect sqrt
		sqrt := math.Sqrt(float64(i))
		if sqrt*sqrt == float64(i) {
			dp[i] = 1
			continue
		}
		for j := 1; j*j <= i; j++ {
			dp[i] = util.Min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
