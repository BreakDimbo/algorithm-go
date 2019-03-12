package tree

// solution 1: recursive
// solution 2: non-recursive

/*
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	preRecur(root, &res)
	return res
}

func preRecur(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	preRecur(node.Left, res)
	preRecur(node.Right, res)
}
*/

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	pNode := root
	for pNode != nil || len(stack) != 0 {
		if pNode != nil {
			res = append(res, pNode.Val)
			stack = append(stack, pNode)
			pNode = pNode.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pNode = node.Right
		}
	}
	return res
}
