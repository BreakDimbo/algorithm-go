package dynamic

func max(slice ...int) int {
	var m int
	for i, v := range slice {
		if i == 0 || v > m {
			m = v
		}
	}
	return m
}

func min(slice ...int) int {
	var m int
	for i, v := range slice {
		if i == 0 || v < m {
			m = v
		}
	}
	return m
}
