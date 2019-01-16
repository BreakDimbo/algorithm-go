package tree

import "math"

/*
	solution 1: 使用中序遍历，判断是否有序
		1. 保存到数组中，最后进行判断
		2. 直接在递归中判断当前节点大于前驱节点
	solution 2: 当前节点大于左子树的最大值，小于右子树的最小值
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* solution 1
var prev *TreeNode

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	prev = nil
	return helper(root)
}

func helper(node *TreeNode) bool {
	if node == nil {
		return true
	}

	if !helper(node.Left) {
		return false
	}
	if prev != nil && prev.Val >= node.Val {
		return false
	}
	prev = node
	return helper(node.Right)
}
*/

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	min, max := math.MinInt64, math.MaxInt64
	return helper(root, min, max)
}

func helper(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if min != math.MinInt64 && node.Val <= min {
		return false
	}
	if max != math.MinInt64 && node.Val >= max {
		return false
	}

	return helper(node.Left, min, node.Val) &&
		helper(node.Right, node.Val, max)
}
