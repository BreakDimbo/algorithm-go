package matrix

/*
	solution 1: use visited
	solution 2: use O(1) space, use first row and first col to remeber
*/

/*
func setZeroes(matrix [][]int) {
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}

	for i := range matrix {
		for j := range matrix[i] {
			if visited[i][j] {
				continue
			}
			visited[i][j] = true
			if matrix[i][j] == 0 {
				for m := 0; m < len(matrix); m++ {
					if matrix[m][j] == 0 {
						continue
					}
					visited[m][j] = true
					matrix[m][j] = 0
				}
				for n := 0; n < len(matrix[i]); n++ {
					if matrix[i][n] == 0 {
						continue
					}
					visited[i][n] = true
					matrix[i][n] = 0
				}
			}
		}
	}
}
*/

func setZeroes(matrix [][]int) {
	colZero := false

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			colZero = true
		}
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for i := range matrix[0] {
			matrix[0][i] = 0
		}
	}

	if colZero {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
}
