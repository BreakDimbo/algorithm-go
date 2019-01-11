package dynamic

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	mini := triangle[len(triangle)-1]
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			mini[j] = min(mini[j], mini[j+1]) + triangle[i][j]
		}
	}
	return mini[0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// DP bottom-to-top

func minTotal(triangle [][]int) int {
	if triangle == nil {
		return 0
	}

	triLen := len(triangle)

	mini := triangle[triLen-1]
	for i := triLen - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			mini[j] = min(mini[j], mini[j+1]) + triangle[i][j]
		}
	}
	return mini[0]
}
