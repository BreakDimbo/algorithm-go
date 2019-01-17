package dynamicp

/*
	solution 1 递归 - * 的处理 - i[len(i):] = [] 不会 index out of bounds
	solution 2 动态规划?
		DP : dp[i][j] -- s[i] - p[j] 时，是否匹配
		状态转移方程： dp[i][j] =
			if s[i] == p[j] then dp[i][j] = dp[i-1][j-1]
			if p[j] == '.' then dp[i][j] = dp[i-1][j-1]
			if p[j] == '*' then
				if p[j-1] != s[i] then dp[i][j] = dp[i][j-2]
				if
*/

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	firstMatch := len(s) != 0 && (p[0] == s[0] || '.' == p[0])
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	}
	return firstMatch && isMatch(s[1:], p[1:])
}
