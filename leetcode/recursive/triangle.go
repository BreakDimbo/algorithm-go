package recursive

import "math"

var min int

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	min = math.MaxInt32

	traverse(triangle, 0, 0, 0)

	return min
}

func traverse(triangle [][]int, i, j, curSum int) {
	if i >= len(triangle) {
		if curSum < min {
			min = curSum
		}
		return
	}
	if j >= len(triangle[i]) {
		return
	}
	curSum += triangle[i][j]
	traverse(triangle, i+1, j, curSum)
	traverse(triangle, i+1, j+1, curSum)
}
