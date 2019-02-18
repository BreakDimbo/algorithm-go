package bdfs

/*
	solution 1: DFS+binary O(nlgn)
	solution 2: start from the top right corner O(m+n)
*/

/*
func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, 0
	for i < len(matrix) && j < len(matrix[i]) {
		if bSearch(matrix, i, j, target) {
			return true
		}
		i++
		j++
	}
	return false
}

func bSearch(m [][]int, i, j, target int) bool {
	// search row
	l, r := j, len(m[i])-1
	for l <= r {
		mid := l + ((r - l) >> 1) // in case of overflow
		if m[i][mid] == target {
			return true
		}
		if m[i][mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	// search col
	low, high := i, len(m)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if m[mid][j] == target {
			return true
		}

		if m[mid][j] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
*/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	col := len(matrix[0]) - 1
	row := 0
	for col >= 0 && row < len(matrix) {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] < target {
			row++
		} else {
			col--
		}
	}
	return false
}
