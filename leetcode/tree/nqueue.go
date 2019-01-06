package tree

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	if n < 1 {
		return nil
	}
	result := [][]int{}
	cols := make(map[int]bool)
	pie := make(map[int]bool)
	na := make(map[int]bool)

	d := &DFS{col: cols, pie: pie, na: na, res: result}
	d.Dfs(n, 0, []int{})
	fmt.Println(d.res)
	return generateResult(d.res, n)
}

func generateResult(r [][]int, n int) [][]string {
	result := make([][]string, 0)
	for _, solution := range r {
		var s string
		var oneSol []string
		for _, colPos := range solution {
			s = strings.Repeat(".", colPos) + "Q" + strings.Repeat(".", n-colPos-1)
			oneSol = append(oneSol, s)
		}
		result = append(result, oneSol)
	}
	return result
}

type DFS struct {
	col map[int]bool
	pie map[int]bool
	na  map[int]bool
	/*
		[[1,3,4,2],
		[2,1,3,4],
		...]
	*/
	res [][]int
}

func (d *DFS) Dfs(n, row int, result []int) {
	// 递归终止条件
	if row >= n {
		d.res = append(d.res, result)
		return
	}

	for col := 0; col < n; col++ {
		// 判断位置是否已经被皇后占用
		if d.col[col] || d.pie[col+row] || d.na[row-col] {
			continue
		}

		// 递的过程，记录皇后所列位置，以及占用的 pie 和 na 的位置
		d.col[col] = true
		d.pie[row+col] = true
		d.na[row-col] = true

		// 开始递归
		d.Dfs(n, row+1, append(result, col))

		// 归的过程恢复现场
		d.col[col] = false
		d.pie[row+col] = false
		d.na[row-col] = false
	}
}
