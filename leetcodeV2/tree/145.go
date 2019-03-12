package tree

/*
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	postRecur(root, &res)
	return res
}

func postRecur(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	postRecur(node.Left, res)
	postRecur(node.Right, res)
	*res = append(*res, node.Val)
}
*/

// reverse the preorder is not good
// general solution
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	pNode := root

	for pNode != nil || len(stack) != 0 {
		if pNode != nil {
			stack = append(stack, pNode)
			res = append([]int{pNode.Val}, res...)
			pNode = pNode.Right
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pNode = node.Left
		}
	}
	return res
}
