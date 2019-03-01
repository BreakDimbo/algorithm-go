package tree

func pathSum(root *TreeNode, sum int) int {
	count := 0
	pathRecursive(root, 0, sum, &count)
	return count
}

func pathRecursive(root *TreeNode, curSum, sum int, count *int) {
	if root == nil {
		return
	}
	if curSum == sum {
		*count++
	}
	if curSum+root.Val > sum {
		pathRecursive(root.Left, root.Val, sum, count)
		pathRecursive(root.Right, root.Val, sum, count)
	} else {
		pathRecursive(root.Left, curSum+root.Val, sum, count)
		pathRecursive(root.Right, curSum+root.Val, sum, count)
	}
}
