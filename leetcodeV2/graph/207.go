package graph

/*
	solution 1: use hash dfs
	solution 2: dfs/bfs topological sort.
*/

/*
func canFinish(numCourses int, prerequisites [][]int) bool {
	prereqMap := make(map[int][]int)

	for i := 0; i < len(prerequisites); i++ {
		if len(prerequisites[i]) == 2 {
			prereqMap[prerequisites[i][0]] = append(prereqMap[prerequisites[i][0]], prerequisites[i][1])
		}
	}

	visited := make(map[int]bool)
	for k := range prereqMap {
		if visited[k] {
			continue
		}
		if detectCycle(prereqMap, make(map[int]bool), visited, k) {
			return false
		}
	}
	return true
}

func detectCycle(preMap map[int][]int, memo, visited map[int]bool, key int) bool {
	if _, ok := preMap[key]; !ok {
		return false
	}
	if memo[key] {
		return true
	}
	memo[key] = true
	visited[key] = true

	for _, newKey := range preMap[key] {
		if detectCycle(preMap, memo, visited, newKey) {
			return true
		}
	}
	memo[key] = false

	return false
}
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make(map[int][]int)
	indegree := make([]int, numCourses)

	for _, pre := range prerequisites {
		adj[pre[1]] = append(adj[pre[1]], pre[0])
		indegree[pre[0]]++
	}

	queue := make([]int, 0)
	// find indegree == 0 node
	for i, n := range indegree {
		if n == 0 {
			queue = append(queue, i)
		}
	}

	cnt := 0
	for ; len(queue) != 0; cnt++ {
		top := queue[0]
		queue = queue[1:]

		for _, crs := range adj[top] {
			indegree[crs]--
			if indegree[crs] == 0 {
				queue = append(queue, crs)
			}
		}
	}
	return cnt == numCourses
}
