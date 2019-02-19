package bdfs

import "github.com/breakD/algorithms/leetcodeV2/util"

/*
	solution 1: dfs
	solution 2: dp + dfs
*/

/*
var (
	max int
	dx  = []int{0, 0, 1, -1}
	dy  = []int{1, -1, 0, 0}
)

func longestIncreasingPath(matrix [][]int) int {
	max := 0
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}
	for i := range matrix {
		for j := range matrix[i] {
			lIPDFS(matrix, visited, i, j, 0)
		}
	}

	return max
}

func lIPDFS(matrix [][]int, visited [][]bool, i, j, curStep int) {
	if curStep > max {
		max = curStep
	}
	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[i]) &&
			!visited[x][y] && matrix[x][y] > matrix[i][j] {
			visited[x][y] = true
			lIPDFS(matrix, visited, x, y, curStep+1)
			visited[x][y] = false
		}
	}
}
*/
var (
	dx = []int{0, 0, 1, -1}
	dy = []int{1, -1, 0, 0}
)

func longestIncreasingPath(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[i]))
	}
	for i := range matrix {
		for j := range matrix[i] {
			lIPDFS(matrix, dp, i, j)
		}
	}

	max := 0
	for i := range matrix {
		for j := range matrix[i] {
			max = util.Max(max, dp[i][j])
		}
	}

	return max
}

func lIPDFS(matrix, dp [][]int, i, j int) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[i]) &&
			matrix[x][y] < matrix[i][j] {
			dp[i][j] = util.Max(dp[i][j], lIPDFS(matrix, dp, x, y))
		}
	}
	return dp[i][j]
}
