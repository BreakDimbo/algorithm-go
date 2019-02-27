package tree

/*
	solution 1: post order
*/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Right = left
	root.Left = right
	return root
}
