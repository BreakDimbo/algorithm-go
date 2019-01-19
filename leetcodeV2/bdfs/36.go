package bdfs

/*
	solution 1 DFS
		剪枝
*/

func isValidSudoku(board [][]byte) bool {
	if len(board) == 0 {
		return false
	}

	for i, cols := range board {
		for j, col := range cols {
			if col == byte('.') {
				continue
			}
			for k := j + 1; k < len(cols); k++ {
				if col == cols[k] {
					return false
				}
			}
			for k := i + 1; k < len(board); k++ {
				if col == board[k][j] {
					return false
				}
			}
			for m := i / 3 * 3; m < (i/3+1)*3; m++ {
				for n := j / 3 * 3; n < (j/3+1)*3; n++ {
					if m == i && n == j {
						continue
					}
					if board[m][n] == board[i][j] {
						return false
					}
				}
			}
		}
	}
	return true
}
