package tree

/*
	solution 1: 递归左右子树，findPorQ
*/

func lowestCommonAncestorP(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestorP(root.Left, p, q)
	right := lowestCommonAncestorP(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left == nil {
		return right
	} else {
		return left
	}
}
