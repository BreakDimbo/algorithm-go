package dynamicp

import "strings"

/*
	solution 1: DFS
	solution 2: DFS+memorize
	solution 3 : DP
*/

/*
func wordBreak(s string, wordDict []string) []string {
	dict := make(map[string]bool)
	for _, w := range wordDict {
		dict[w] = true
	}
	res := make([]string, 0)

	helper(s, dict, "", &res)
	return res
}

func helper(s string, dict map[string]bool, curStr string, res *[]string) {
	if len(s) == 0 {
		curStr = strings.TrimSuffix(curStr, " ")
		*res = append(*res, curStr)
		return
	}

	for i := 1; i <= len(s); i++ {
		if dict[s[0:i]] {
			helper(s[i:], dict, curStr+s[0:i]+" ", res)
		}
	}
}
*/

func wordBreak(s string, wordDict []string) []string {
	dict := make(map[string]bool)
	for _, w := range wordDict {
		dict[w] = true
	}
	memo := make(map[string][]string)

	return helper(s, dict, memo)
}

func helper(s string, dict map[string]bool, memo map[string][]string) []string {
	if value, ok := memo[s]; ok {
		return value
	}
	res := make([]string, 0)
	if len(s) == 0 {
		res = append(res, "") // 这里要 append，否则 restResult 永远为空
		return res
	}

	for word := range dict {
		if strings.HasPrefix(s, word) {
			restResult := helper(s[len(word):], dict, memo)
			for _, item := range restResult {
				var sep string
				if len(item) != 0 {
					sep = " "
				}
				res = append(res, word+sep+item)
			}
		}
	}
	memo[s] = res
	return res
}
