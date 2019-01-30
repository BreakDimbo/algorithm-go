package bdfs

/*
	solution 1: DP
*/

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 0
	if s[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}

	for i := 2; i <= len(s); i++ {
		if (s[i-1]-'0')+(s[i-2]-'0')*10 <= 26 && (s[i-1]-'0')+(s[i-2]-'0')*10 > 0 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(s)]
}
