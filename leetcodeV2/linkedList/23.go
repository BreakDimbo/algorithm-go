package linkedList

/*
	solution 1: 遍历
	solution 2: 遍历所有节点，把值取出来放进数组，排序后放进新链表
	solution 3: 使用两个优先队列
	solution 4: 分治
*/

/*
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	dummy := &ListNode{}
	head := dummy
	index := 0

	for {
		minNode := &ListNode{Val: math.MaxInt32}
		for i, list := range lists {
			if list != nil && list.Val <= minNode.Val {
				minNode = list
				index = i
			}
		}
		if minNode.Val == math.MaxInt32 {
			break
		}
		dummy.Next = minNode
		lists[index] = lists[index].Next
		dummy = dummy.Next
	}

	return head.Next
}
*/

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return dividConquer(lists, 0, len(lists)-1)
}

func dividConquer(l []*ListNode, start, end int) *ListNode {
	if start > end {
		return nil
	}
	if start == end {
		return l[start]
	}

	mid := start + ((end - start) >> 1)
	left := dividConquer(l, start, mid)
	right := dividConquer(l, mid+1, end)

	return mergeLists(left, right)
}

func mergeLists(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	} else {
		cur.Next = right
	}
	return dummy.Next
}
