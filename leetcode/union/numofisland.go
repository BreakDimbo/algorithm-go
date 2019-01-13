package union

// vector array
var (
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, 1, -1}
)

// Flood Fill
func numIslandS(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	sum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum += floodFillDFS(i, j, grid)
		}
	}
	return sum
}

func floodFillDFS(i, j int, grid [][]byte) int {
	if !isValid(i, j, grid) {
		return 0
	}

	grid[i][j] = byte('0')

	for k := 0; k < 4; k++ {
		floodFillDFS(i+dx[k], j+dy[k], grid)
	}
	return 1
}

func isValid(i, j int, grid [][]byte) bool {
	// 越界
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return false
	}
	// 非岛屿
	if grid[i][j] != byte('1') {
		return false
	}
	return true
}

// Union
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	u := NewUnionFind(grid)

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == byte('0') {
				continue
			}
			for k := 0; k < 4; k++ {
				x, y := i+dx[k], j+dy[k]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == byte('1') {
					u.Union(i*n+j, x*n+y)
				}
			}
		}
	}
	return u.count
}
