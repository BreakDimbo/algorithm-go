package bdfs

func generateParenthesis(n int) []string {
	// 保证）的数量永远小于等于（
	res := make([]string, 0)
	parenthDfs(n, 0, 0, "", &res)
	return res
}

func parenthDfs(n, lc, rc int, curStr string, res *[]string) {
	if lc == n && rc == n {
		*res = append(*res, curStr)
		return
	}

	if lc < n {
		parenthDfs(n, lc+1, rc, curStr+"(", res)
	}
	if rc < n && rc < lc {
		parenthDfs(n, lc, rc+1, curStr+")", res)
	}
}

/*
func solveNQueens(n int) [][]string {
	sols := make([][]int, 0)
	col, pie, na := make([]int, n), make([]int, 2*n), make([]int, 2*n)
	var dfsQueue func(n, row int, curSol []int)
	dfsQueue = func(n, row int, curSol []int) {
		if row >= n {
			dst := make([]int, n)
			copy(dst, curSol)
			sols = append(sols, curSol)
		}

		for i := 0; i < n; i++ {
			if col[i] == 1 || pie[i+row] == 1 || na[i-row+n] == 1 {
				continue
			}
			col[i] = 1
			pie[i+row] = 1
			na[i-row+n] = 1
			dfsQueue(n, row+1, append(curSol, i))
			col[i] = 0
			pie[i+row] = 0
			na[i-row+n] = 0
		}
	}
	dfsQueue(n, 0, make([]int, n))
	return printRes(sols)
}
*/

func solveNQueens(n int) [][]string {
	sols := make([][]int, 0)
	col, pie, na := make([]int, n), make([]int, 2*n), make([]int, 2*n)
	var dfsQueue func(n, row int, curSol []int)
	dfsQueue = func(n, row int, curSol []int) {
		if row >= n {
			dst := make([]int, n)
			copy(dst, curSol)
			sols = append(sols, curSol)
		}

		for i := 0; i < n; i++ {
			if col[i] == 1 || pie[i+row] == 1 || na[i-row+n] == 1 {
				continue
			}
			col[i] = 1
			pie[i+row] = 1
			na[i-row+n] = 1
			dfsQueue(n, row+1, append(curSol, i))
			col[i] = 0
			pie[i+row] = 0
			na[i-row+n] = 0
		}
	}
	dfsQueue(n, 0, make([]int, n))
	return printRes(sols)
}
