package matrix

/*
	solution 1: simulation
		use direction array
	solution 2: layer-by-layer
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	ans := []int{}
	m, n := len(matrix), len(matrix[0])
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}

	r, c, di := 0, 0, 0
	for i := 0; i < m*n; i++ {
		ans = append(ans, matrix[r][c])
		seen[r][c] = true

		rr, cc := r+dr[di], c+dc[di]
		if rr >= 0 && rr < m && cc >= 0 && cc < n && !seen[rr][cc] {
			r = rr
			c = cc
		} else {
			di = (di + 1) % 4
			r += dr[di]
			c += dc[di]
		}
	}
	return ans
}
