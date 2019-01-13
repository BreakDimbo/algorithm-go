package union

type Union struct {
	count  int   // 独立集合的个数
	rank   []int // 并集数的高度
	parent []int // 节点 i 的父节点
}

func NewUnionFind(grid [][]byte) *Union {
	m, n := len(grid), len(grid[0])
	u := &Union{rank: make([]int, m*n), parent: make([]int, m*n)}
	for i := 0; i < m*n; i++ {
		u.rank[i] = 0
		u.parent[i] = -1
	}
	for i := range grid {
		for j := range grid[i] {
			// 初始化，将二维数组转换为一维数组
			if grid[i][j] == byte('1') {
				u.parent[i*n+j] = i*n + j
				u.count++
			}
		}
	}
	return u
}

func (u *Union) Find(i int) int {
	for u.parent[i] != i {
		i = u.parent[i]
	}
	return u.parent[i]
}

func (u *Union) Union(x, y int) {
	rootx := u.Find(x)
	rooty := u.Find(y)
	// 注意判断只有在老大不相同时才合并
	if rootx != rooty {
		if u.rank[rootx] > u.rank[rooty] {
			u.parent[rooty] = rootx
		} else if u.rank[rooty] > u.rank[rootx] {
			u.parent[rootx] = rooty
		} else {
			u.parent[rooty] = rootx
			u.rank[rootx]++
		}
		u.count--
	}
}
