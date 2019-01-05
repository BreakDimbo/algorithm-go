package tree

import (
	"math"
)

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	leftMin := minDepth(root.Left)
	rightMin := minDepth(root.Right)

	return 1 + int(math.Min(float64(leftMin), float64(rightMin)))
}
