package tree

/*
	solution 1: preOrder is root, inorder find left and right
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	rootIndex := 0
	for rootVal != inorder[rootIndex] {
		rootIndex++
	}

	left := buildTree(preorder[1:rootIndex+1], inorder[0:rootIndex])
	right := buildTree(preorder[rootIndex+1:len(preorder)], inorder[rootIndex+1:len(inorder)])
	return &TreeNode{
		Val:   rootVal,
		Left:  left,
		Right: right,
	}
}
