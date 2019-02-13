package tree

import (
	"math"

	"github.com/breakD/algorithms/leetcodeV2/util"
)

/*
	solution 1: DP
	solution 2: recursive
*/

var maxValue int

func maxPathSum(root *TreeNode) int {
	maxValue = math.MinInt32
	maxDown(root)
	return maxValue
}

func maxDown(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := util.Max(0, maxDown(root.Left))
	right := util.Max(0, maxDown(root.Right))
	maxValue = util.Max(maxValue, left+right+root.Val)
	return util.Max(left, right) + root.Val
}
