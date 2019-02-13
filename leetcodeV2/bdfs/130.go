package bdfs

/*
	solution 1: dfs
*/

func solveF(board [][]byte) {
	for i := range board {
		for j := range board[i] {
			if i == 0 || i == len(board)-1 || j == 0 || j == len(board[i])-1 {
				if board[i][j] == 'O' {
					flip(i, j, board)
				}
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '*' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func flip(i, j int, board [][]byte) {
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[i])-1 {
		return
	}

	if board[i][j] == 'O' {
		board[i][j] = '*'
		for k := 0; k < 4; k++ {
			flip(dx[k]+i, dy[k]+j, board)
		}
	}
}
