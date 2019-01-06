package tree

import "testing"

func Test_isValidBST(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
	}
	tree.Left = &TreeNode{
		Val: 1,
	}
	// isValidBST(tree)
}
