package bdfs

/*
	solution 1: dfs + visited
*/

var (
	dx = []int{0, 0, -1, 1}
	dy = []int{1, -1, 0, 0}
)

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}

	visited := make([][]int, len(board))
	for k := range visited {
		visited[k] = make([]int, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if wordDfs(board, i, j, word, visited) {
				return true
			}
		}
	}
	return false
}

func wordDfs(board [][]byte, i, j int, word string, visited [][]int) bool {
	if len(word) == 0 {
		return true
	}
	if i < 0 || i >= len(board) ||
		j < 0 || j >= len(board[i]) ||
		board[i][j] != word[0] || visited[i][j] == 1 {
		return false
	}

	visited[i][j] = 1
	for k := 0; k < 4; k++ {
		x, y := i+dx[k], j+dy[k]
		if wordDfs(board, x, y, word[1:], visited) {
			return true
		}
	}
	visited[i][j] = 0
	return false
}
