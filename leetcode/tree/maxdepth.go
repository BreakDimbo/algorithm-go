package tree

import "math"

//* Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
}


// 问题：如何知道当前在哪一层？
func maxDepthBfs(root *TreeNode) int {
	queue := []*TreeNode{}
	visited := make(map[*TreeNode]bool)
	queue = append(queue, root)

	for queue != nil {
		node := queue[0]
		queue = queue[1:]
		if visited[node] {
			continue
		}
		visited[node] = true

		// process data

		nodes := generateChildNodes(root)
		queue = append(queue, nodes...)
	}
	return 0
}

func generateChildNodes(node *TreeNode) []*TreeNode {
	if node.Left != nil && node.Right != nil {
		return []*TreeNode{node.Left, node.Right}
	} else if node.Left != nil {
		return []*TreeNode{node.Left}
	} else if node.Right != nil {
		return []*TreeNode{node.Right}
	}
	return nil
}
