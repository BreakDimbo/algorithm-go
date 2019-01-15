package util

func Max(slice ...int) int {
	var m int
	for i, v := range slice {
		if i == 0 || v > m {
			m = v
		}
	}
	return m
}

func Min(slice ...int) int {
	var m int
	for i, v := range slice {
		if i == 0 || v < m {
			m = v
		}
	}
	return m
}
