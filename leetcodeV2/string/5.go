package string

/*
	solution 1: brute force 两层循环查找
	solution 2: 翻转字符串，找最长 common substring, 但是找到后需要比较 indices 与 翻转后的 S` 的 indices 是否相同
	solution 3: DP：P(i,j) 表示 [i,j] 是否是 palindrome
		状态转移方程：P(i,j) = P(i+1,j-1) && s[i] == s[j]
		注意是从后向前循环，因为如果从前向后循环，就不知道 dp[i+1][j-1] true of false, 因为 i 是由 i+1 推出来的
	solution 4: Expand Around Center
*/

/*
func longestPalindrome(s string) string {
	n := len(s)
	ans := ""
	// init state
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
				if j-i+1 > len(ans) {
					ans = s[i : j+1]
				}
			}
		}
	}
	return ans
}
*/

func longestPalindrome(s string) string {
	ans := ""
	exp := func(s string, left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		if right-left-1 > len(ans) {
			ans = s[left+1 : right]
		}
	}
	for i := range s {
		exp(s, i, i)
		exp(s, i, i+1)
	}
	return ans
}
