package tree

import "math"

var pre *TreeNode

func isValidBST(root *TreeNode) bool {
	pre = nil
	return helper(root)
}

func helper(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !helper(root.Left) {
		return false
	}
	if pre != nil && pre.Val >= root.Val {
		return false
	}
	pre = root
	return helper(root.Right)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return right
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	for root != nil {
		if root.Val > q.Val && root.Val > p.Val {
			root = root.Left
		} else if root.Val < q.Val && root.Val < p.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return root
}

func maxDepth(root *TreeNode) int {
	maxDepth := 0
	depthDfs(root, 0, &maxDepth)
	return maxDepth
}

func depthDfs(node *TreeNode, depth int, maxDepth *int) {
	if node == nil {
		if depth > *maxDepth {
			*maxDepth = depth
		}
		return
	}

	depthDfs(node.Left, depth+1, maxDepth)
	depthDfs(node.Right, depth+1, maxDepth)
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minDepth := math.MaxInt32
	minDepthDfs(root, 0, &minDepth)
	return minDepth
}

func minDepthDfs(node *TreeNode, depth int, minDepth *int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil && depth < *minDepth {
		*minDepth = depth
		return
	}
	minDepthDfs(node.Left, depth+1, minDepth)
	minDepthDfs(node.Right, depth+1, minDepth)
}
