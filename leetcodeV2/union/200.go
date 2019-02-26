package union

/*
	solution 1: flood Fill
	solution 2: union
*/

/*
var (
	dx = []int{0, 0, 1, -1}
	dy = []int{1, -1, 0, 0}
)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == byte('1') {
				DFSMarking(grid, i, j)
				count++
			}
		}
	}
	return count
}

func DFSMarking(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid) || grid[i][j] != byte('1') {
		return
	}
	grid[i][j] = '0'
	for k := 0; k < 4; k++ {
		x := dx[k] + i
		y := dy[k] + j
		DFSMarking(grid, x, y)
	}
}
*/

type UnionFU struct {
	roots []int
}

func NewUnion(n int) *UnionFU {
	return &UnionFU{
		make([]int, n),
	}
}

func (u *UnionFU) FindRood(i int) int {
	root := i
	for root != u.roots[root] {
		root = u.roots[root]
	}

	// 路径压缩
	for i != root {
		u.roots[i], i = root, u.roots[i]
	}
	return root
}

func (u *UnionFU) Union(a, b int) {
	rootA := u.FindRood(a)
	rootB := u.FindRood(b)
	u.roots[rootA] = rootB
}
