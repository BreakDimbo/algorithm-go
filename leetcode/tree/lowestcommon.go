package tree

func lowesCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	} else if right == nil {
		return left
	}
	return root
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if q.Val > root.Val && p.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if q.Val < root.Val && p.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}
