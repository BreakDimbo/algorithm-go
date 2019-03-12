package tree

/*
	solution 1: recursive
		inorder - pre -> in -> after
	solution 2: interate
		use stack
*/

/*
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
*/

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	pNode := root

	for pNode != nil || len(stack) != 0 {
		if pNode != nil {
			stack = append(stack, pNode)
			pNode = pNode.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
			pNode = node.Right
		}
	}
	return res
}
