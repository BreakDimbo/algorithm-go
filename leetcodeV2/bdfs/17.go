package bdfs

/*
	solution 1: DFS, 7&9
*/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	ans := []string{}
	mapping := []string{"", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	var dfs func(d, cs string)
	dfs = func(digits, curStr string) {
		if len(digits) == 0 {
			ans = append(ans, curStr)
			return
		}
		d := digits[0] - '1'
		for i := mapping[d][0]; i <= mapping[d][len(mapping[d])-1]; i++ {
			dfs(digits[1:], curStr+string(i))
		}
	}

	dfs(digits, "")
	return ans
}
