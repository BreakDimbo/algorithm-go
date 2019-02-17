package tree

import "math"

/*
	solution 1: 中序遍历
*/

func kthSmallest(root *TreeNode, k int) int {
	min := math.MaxInt32
	kthDfs(root, &k, &min)
	return min
}

func kthDfs(root *TreeNode, k, min *int) {
	if root == nil {
		return
	}
	kthDfs(root.Left, k, min)
	*k--
	if *k == 0 {
		*min = root.Val
		return
	}
	kthDfs(root.Right, k, min)
}
