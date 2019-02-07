package tree

import "container/list"

/*
	solution 1: DFS
			level & 1 == 0
*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	container := make([]*list.List, 0)
	result := make([][]int, 0)

	zigLevelHelper(&container, root, 0)

	for _, l := range container {
		arr := make([]int, 0)
		for e := l.Front(); e != nil; e = e.Next() {
			arr = append(arr, e.Value.(int))
		}
		result = append(result, arr)
	}
	return result
}

func zigLevelHelper(con *[]*list.List, root *TreeNode, height int) {
	if root == nil {
		return
	}
	if height >= len(*con) {
		*con = append(*con, list.New())
	}
	if height&1 == 0 {
		(*con)[height].PushBack(root.Val)
	} else {
		(*con)[height].PushFront(root.Val)
	}

	zigLevelHelper(con, root.Left, height+1)
	zigLevelHelper(con, root.Right, height+1)
}
