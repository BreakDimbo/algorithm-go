package tree

import "container/list"

func convertBST(root *TreeNode) *TreeNode {
	stack := list.New()
	pNode := root
	sum := 0

	for pNode != nil || stack.Len() > 0 {
		if pNode != nil {
			stack.PushBack(pNode)
			pNode = pNode.Right
		} else {
			node := stack.Remove(stack.Back()).(*TreeNode)
			sum += node.Val
			node.Val = sum
			pNode = node.Left
		}
	}
	return root
}
