package tree

/*
	solution 1: recursive
		seprate the tree
	solution 2: iterative
		BFS
*/

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}

	return (t1.Val == t2.Val &&
		isMirror(t1.Left, t2.Right) &&
		isMirror(t1.Right, t2.Left))
}
