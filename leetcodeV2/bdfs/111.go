package bdfs

import "github.com/breakD/algorithms/leetcodeV2/util"

/*
	solution 1: DFS
	solution 2: BFS
*/

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return util.Min(minDepth(root.Left), minDepth(root.Right)) + 1
}
