package string

/*
	solution 1: dfs
*/

func partition(s string) [][]string {
	res := make([][]string, 0)
	if len(s) == 0 {
		return nil
	}

	dfs(s, nil, &res)

	return res
}

func dfs(s string, curSet []string, res *[][]string) {
	if len(s) == 0 {
		set := make([]string, len(curSet))
		copy(set, curSet)
		*res = append(*res, set)
		return
	}

	for j := 1; j <= len(s); j++ {
		if isPalin(s[0:j]) {
			dfs(s[j:], append(curSet, s[0:j]), res)
		}
	}

}

func isPalin(s string) bool {
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
