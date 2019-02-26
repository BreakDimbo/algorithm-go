package union

/*
	solution 1: union
	solution 2: Flood fill
*/

func findCircleNum(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	visited := make([]bool, len(M))
	count := 0
	for i := range M {
		if !visited[i] {
			DFSMark(M, i, visited)
			count++
		}
	}
	return count
}

func DFSMark(M [][]int, person int, visited []bool) {
	for other := 0; other < len(M); other++ {
		if !visited[other] && M[person][other] == 1 {
			visited[other] = true
			DFSMark(M, other, visited)
		}
	}
}
