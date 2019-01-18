package bdfs

// func dfs() {

// }

func bfs(root *TreeNode) {
	if root == nil {
		return
	}

	visited := make(map[*TreeNode]bool)
	queue := []*TreeNode{}
	queue = append(queue, root)

	for queue != nil {
		pop := queue[0]
		queue = queue[1:]
		if visited[pop] {
			continue
		}
		visited[pop] = true
		// process data

		// get children and append to queue
		queue = append(queue, []*TreeNode{pop.Right, pop.Left}...)
	}
}
