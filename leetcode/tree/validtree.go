package tree

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	arr := traverseInorder(root)

	// 判断重复
	for i := 1; i < len(arr); i++ {
		if arr[i-1] == arr[i] {
			return false
		}
	}

	return sort.IntsAreSorted(arr)
}

func traverseInorder(n *TreeNode) []int {
	if n == nil {
		return []int{}
	}
	l := traverseInorder(n.Left)
	m := n.Val
	r := traverseInorder(n.Right)
	l = append(l, m)
	return append(l, r...)
}
*/

/*
var prev *TreeNode

func isValidBST(root *TreeNode) bool {
	prev = nil
	return helper(root)
}

func helper(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if !helper(root.Left) {
		return false
	}
	if prev != nil && prev.Val >= root.Val {
		return false
	}
	prev = root
	return helper(root.Right)
}
*/

// func isValidBST(root *TreeNode) bool {

// }
