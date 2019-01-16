package tree

/*
	solution 1: 利用二叉搜索树的性质，
	判断当前节点 r 的值与 p q 的值，找到第一次 !(r > q && r > p || r < q && r < p)
*/

func lowestCommonAncestorR(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val > q.Val && root.Val > p.Val {
		return lowestCommonAncestorR(root.Left, p, q)
	} else if root.Val < q.Val && root.Val < p.Val {
		return lowestCommonAncestorR(root.Right, p, q)
	}
	return root
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root == nil {
			return root
		}
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return root
}
