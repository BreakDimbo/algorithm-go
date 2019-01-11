package bit

var count int

func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}
	count = 0
	dfs(n, 0, 0, 0, 0)
	return count
}

func dfs(n, row, col, pie, na int) {
	// 此时 col, pie, na 中 1 代表占位，0 代表空位
	// 如果 row 超过 n 说明找到了一个解
	if row >= n {
		count++
		return
	}

	// 找出当前行所有空位,此时 col, pie, na 中 1 代表空位，0 代表占位
	bits := (^(col | pie | na)) &
		((1 << uint(n)) - 1)

	for bits != 0 {
		// 找出最后一个 1 所在的位置，其他位均变为 0，执行占位，此时 1 表示占位
		p := bits & -bits
		// 执行递归，更新 col, pie, na,注意 pie 和 na 要分别左移一位，右移一位
		dfs(n, row+1, col|p, (pie|p)<<1, (na|p)>>1)

		// 消掉最后一位1，继续循环
		// bits = bits & (bits - 1)
		bits &= bits - 1
	}
}
