package bit

/*
	solution 1:
		00 : dead(next) <- dead(current)
		01 : dead(next) <- live(current)
		10 : live(next) <- dead(current)
		11 : live(next) <- live(current)

		result >> 1 will get the 2nd state
		check result & 1 get the current state
*/

var (
	dx = []int{0, 0, 1, -1, -1, 1, -1, 1}
	dy = []int{-1, 1, 0, 0, 1, 1, -1, -1}
)

func gameOfLife(board [][]int) {
	if len(board) == 0 {
		return
	}

	for i := range board {
		for j := range board[i] {
			lives := liveCntOfNeighbors(board, i, j)
			// we just care about the next state will be live
			if board[i][j] == 1 && lives >= 2 && lives <= 3 {
				board[i][j] = 3
			}
			if board[i][j] == 0 && lives == 3 {
				board[i][j] = 2
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] >>= 1
		}
	}
}

func liveCntOfNeighbors(board [][]int, i, j int) int {
	lives := 0
	for k := 0; k < 8; k++ {
		x := i + dx[k]
		y := j + dy[k]
		if x >= 0 && y >= 0 && x < len(board) && y < len(board[i]) {
			lives += board[x][y] & 1
		}
	}
	return lives
}
