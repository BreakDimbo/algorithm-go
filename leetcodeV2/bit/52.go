package bit

/*
	solution 1: DFS
	solution 2: bits
*/

func totalNQueens(n int) int {
	count := 0
	// availeble position stands by 0
	queueDfs(n, 0, 0, 0, 0, &count)
	return count
}

func queueDfs(n, row, col, pie, na int, count *int) {
	if row >= n {
		*count++
		return
	}

	// find the available position, stands by 1
	bits := (^(col | pie | na)) & ((1 << uint(n)) - 1)

	for bits != 0 {
		// find the last available 1, now 1 stands for unavalaible
		p := bits & -bits
		queueDfs(n, row+1, col|p, (pie|p)<<1, (na|p)>>1, count)

		// change the last 1 to 0
		bits &= bits - 1
	}
}
