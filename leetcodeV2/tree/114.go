package tree

/*
	solution 1: backtrace + post order
*/

func flatten(root *TreeNode) {

	var flattenR func(root *TreeNode)
	var cur *TreeNode

	flattenR = func(root *TreeNode) {
		if root == nil {
			return
		}
		flattenR(root.Right)
		flattenR(root.Left)
		root.Right = cur
		root.Left = nil
		cur = root
	}

	flattenR(root)
}
