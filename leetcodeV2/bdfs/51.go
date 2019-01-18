package bdfs

import "strings"

func solveNQueens(n int) [][]string {
	if n < 1 {
		return nil
	}
	result := make([][]int, 0)
	cols, pie, na := make([]int, n), make([]int, 2*n), make([]int, 2*n)
	var dfs func(n, row int, curState []int)

	dfs = func(n, row int, curState []int) {
		if row >= n {
			dst := make([]int, len(curState))
			copy(dst, curState) // 注意这里要进行一份 copy，否则 result 是指针，递归外部可能修改该值
			result = append(result, dst)
			return
		}

		for col := 0; col < n; col++ {
			if cols[col] == 1 || pie[col+row] == 1 || na[row-col+n] == 1 {
				continue
			}

			cols[col] = 1
			pie[col+row] = 1
			na[row-col+n] = 1

			dfs(n, row+1, append(curState, col))

			cols[col] = 0
			pie[col+row] = 0
			na[row-col+n] = 0
		}
	}
	
	dfs(n, 0, make([]int, 0))

	return generateResult(result, n)
}

func generateResult(res [][]int, n int) [][]string {
	solution := [][]string{}

	for _, v := range res {
		oneSolu := []string{}
		for _, pos := range v {
			s := strings.Repeat(".", pos) + "Q" + strings.Repeat(".", n-pos-1)
			oneSolu = append(oneSolu, s)
		}
		solution = append(solution, oneSolu)
	}
	return solution
}
