package bdfs

/*
	solution 1: Dynamic programming
*/

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[n] = 1
	if s[n-1] == '0' {
		dp[n-1] = 0
	} else {
		dp[n-1] = 1
	}

	for i := n - 2; i >= 0; i-- {
		if s[i] == '0' {
			continue
		}
		if (s[i+1]-'0')+(s[i]-'0')*10 > 26 {
			dp[i] = dp[i+1]
		} else {
			dp[i] = dp[i+1] + dp[i+2]
		}
	}
	return dp[0]
}
