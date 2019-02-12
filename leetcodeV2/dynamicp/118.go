package dynamicp

/*
	solution 1: iterate
*/

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	if numRows == 1 {
		return [][]int{[]int{1}}
	}

	result := make([][]int, numRows)
	result[0] = []int{1}
	result[1] = []int{1, 1}

	for i := 3; i <= numRows; i++ {
		row := make([]int, i)
		row[0], row[i-1] = 1, 1
		for k := 0; k < i-2; k++ {
			row[k+1] = result[i-2][k] + result[i-2][k+1]
		}
		result[i-1] = row
	}
	return result
}
