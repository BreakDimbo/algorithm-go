package bdfs

/*
	solution 1: 遍历所有可能，找出有效组合
	solution 2: 左括号的个数永远 >= 右括号的个数，DFS
*/

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var pair func(n, i, j int, s string)
	pair = func(n, i, j int, s string) {
		if i > n || j > n || j > i {
			return
		}

		pair(n, i+1, j, s+"(")
		pair(n, i, j+1, s+")")

		if j == i && i == n {
			res = append(res, s)
		}
	}

	pair(n, 0, 0, "")
	return res
}

func pair(n, i, j int, s string) {
	if i > n || j > n || j > i {
		return
	}

	pair(n, i+1, j, s+"(")
	pair(n, i, j+1, s+")")

	if j == i && i == n {
		// res = append(res, s)
	}
}

func pair2(n, i, j int, s string) {
	if i == n && j == n {
		// res = append(res, s)
	}
	if i < n {
		pair(n, i+1, j, s+"(")
	}
	if j < n && j < i {
		pair(n, i, j+1, s+")")
	}
}
