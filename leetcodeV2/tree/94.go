package tree

/*
	solution 1: recursive
		inorder - pre -> in -> after
	solution 2: interate
		use stack
*/

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		result = append(result, root.Val)
		inorder(root.Right)
	}

	inorder(root)
	return result
}
