package matrix

/*
	solution 1: heap,每次 pop 出最小元素后，用相同列的下一个元素替换
	solution 2: binary， search space is range, not index
*/

func kthSmallest(matrix [][]int, k int) int {
	lo, hi := matrix[0][0], matrix[len(matrix)-1][len(matrix)-1]

	for lo < hi {
		mid := lo + ((hi - lo) >> 1)
		count := 0
		for i := range matrix {
			j := len(matrix[i]) - 1
			for ; j >= 0 && matrix[i][j] > mid; j-- {
			}
			count += (j + 1)
		}
		if count < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
