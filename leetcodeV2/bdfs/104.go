package bdfs

import "github.com/breakD/algorithms/leetcodeV2/util"

/*
	solution 1: DFS - 递归或者使用栈
	solution 2: BFS
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return util.Max(maxDepth(root.Right), maxDepth(root.Left)) + 1
}
