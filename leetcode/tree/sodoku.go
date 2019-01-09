package tree

func solveSudoku(board [][]byte) {
	if len(board) == 0 {
		return
	}
	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == acc2b(".") { // 注意判断只有为空的格子才向里面填子
				for c := acc2b("1"); c <= acc2b("9"); c++ {
					if isValid(board, i, j, c) {
						board[i][j] = c
						if solve(board) {
							return true
						} else {
							board[i][j] = acc2b(".")
						}
					}
				}
				return false // 注意只有在 1～9 都尝试完，才返回 false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, c byte) bool {
	for i := 0; i < len(board); i++ {
		if board[row][i] != acc2b(".") && board[row][i] == c {
			return false
		}
		if board[i][col] != acc2b(".") && board[i][col] == c {
			return false
		}
		// 注意这里对 3*3 格子判断对写法
		if board[3*(row/3)+i/3][3*(col/3)+i%3] != acc2b(".") && board[3*(row/3)+i/3][3*(col/3)+i%3] == c {
			return false
		}
	}
	return true
}

func acc2b(i string) byte {
	return []byte(i)[0]
}
