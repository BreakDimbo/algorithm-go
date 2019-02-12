package tree

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	rootIndex := len(nums) / 2
	rootVal := nums[rootIndex]

	left := sortedArrayToBST(nums[0:rootIndex])
	right := sortedArrayToBST(nums[rootIndex+1 : len(nums)])
	return &TreeNode{
		Val:   rootVal,
		Left:  left,
		Right: right,
	}
}
