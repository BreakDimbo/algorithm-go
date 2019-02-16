package bdfs

/*
	solution 1: BFS topological sort
*/

func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := make(map[int][]int)
	indegree := make([]int, numCourses)
	res := make([]int, 0)

	for _, pre := range prerequisites {
		adj[pre[1]] = append(adj[pre[1]], pre[0])
		indegree[pre[0]]++
	}

	queue := make([]int, 0)
	for i, degree := range indegree {
		if degree == 0 {
			queue = append(queue, i)
			indegree[i]--
		}
	}

	for len(queue) != 0 {
		top := queue[0]
		res = append(res, top)
		queue = queue[1:]

		for _, crs := range adj[top] {
			indegree[crs]--
			if indegree[crs] == 0 {
				queue = append(queue, crs)
			}
		}
	}

	if len(res) == numCourses {
		return res
	}
	return nil
}
