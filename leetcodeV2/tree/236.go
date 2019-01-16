package tree

/*
	solution 1: 递归左右子树，findPorQ
	solution 2: 判断当前节点 r 的值与 p q 的值，找到第一次 !(r > q && r > p || r < q && r < p) - 只能用在二叉搜索树中
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left == nil {
		return right
	} else {
		return left
	}
}
