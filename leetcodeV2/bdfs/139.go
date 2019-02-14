package bdfs

/*
	solution 1: dfs
	solution 2: dp
*/

/*
func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, w := range wordDict {
		dict[w] = true
	}

	return canBreak(s, dict)
}

func canBreak(s string, dict map[string]bool) bool {
	if len(s) == 0 {
		return true
	}

	for i := 1; i <= len(s); i++ {
		if dict[s[0:i]] {
			if canBreak(s[i:], dict) {
				return true
			}
		}
	}
	return false
}
*/

func wordBreak(s string, wordDict []string) bool {
	// f[i] stands for if s[0:i] can be broken
	f := make([]bool, len(s))
	f[0] = true

	dict := make(map[string]bool)
	for _, w := range wordDict {
		dict[w] = true
	}

	for i := 1; i <= len(s); i++ {
		for j := 0; j <= i; j++ {
			if f[j] && dict[s[j:i]] {
				f[i] = true
				break
			}
		}
	}
	return f[len(s)-1]
}
