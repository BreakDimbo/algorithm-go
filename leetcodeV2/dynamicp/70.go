package dynamicp

/*
	solution 1: dp
	solution 2: recursive
*/

func climbStairs(n int) int {
	// dp[i] = dp[i-1]+dp[i-2]
	cur, pre := 1, 1
	for i := 2; i <= n; i++ {
		cur, pre = cur+pre, cur
	}
	return cur
}
