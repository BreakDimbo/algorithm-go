package tree

/*
	solution 1: DFS
	solution 2: recursive
*/

func levelOrder(root *TreeNode) [][]int {
	queue := make([][]int, 0)
	levelHelper(&queue, root, 0)
	return queue
}

func levelHelper(queue *[][]int, root *TreeNode, height int) {
	if root == nil {
		return
	}
	if height >= len(*queue) {
		*queue = append(*queue, make([]int, 0))
	}
	(*queue)[height] = append((*queue)[height], root.Val)
	levelHelper(queue, root.Left, height+1)
	levelHelper(queue, root.Right, height+1)
}
