package tree

import "github.com/breakD/algorithms/leetcodeV2/util"

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 1
	traverse(root, &ans)
	return ans - 1
}

func traverse(node *TreeNode, ans *int) int {
	if node == nil {
		return 0
	}
	l := traverse(node.Left, ans)
	r := traverse(node.Right, ans)
	*ans = util.Max(*ans, l+r+1)
	return util.Max(l, r) + 1
}
