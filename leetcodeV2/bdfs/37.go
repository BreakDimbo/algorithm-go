package bdfs

/*
	solution 1: DFS
*/
func solveSudoku(board [][]byte) {
	if len(board) == 0 {
		return
	}

	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != byte('.') {
				continue
			}
			for k := '1'; k <= '9'; k++ {
				if isValid(board, i, j, k) {
					board[i][j] = byte(k)
					if solve(board) {
						return true
					} else {
						board[i][j] = byte('.')
					}
				}
			}
			return false
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, char rune) bool {
	for i := 0; i < len(board); i++ {
		if board[row][i] != byte('.') && board[row][i] == byte(char) {
			return false
		}
		if board[i][col] != byte('.') && board[i][col] == byte(char) {
			return false
		}
		if board[3*(row/3)+i/3][3*(col/3)+i%3] != byte('.') && board[3*(row/3)+i/3][3*(col/3)+i%3] == byte(char) {
			return false
		}
	}
	return true
}
