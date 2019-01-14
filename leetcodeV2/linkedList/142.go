package linkedList

/*
	solution 1: 使用 map，返回第一次遇到的重复元素
	solution 2: 双指针：2k-k = nr
*/

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	visited := make(map[*ListNode]bool)
	for head != nil {
		if visited[head] {
			return head
		}
		visited[head] = true
		head = head.Next
	}
	return nil
}

func detectCycleV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// 注意：fast 一定要从 head 开始，否则无法满足 fast 走过的路程是 slow 的两倍
	slow, fast := head, head

	// find the meet point
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			for head != fast {
				head = head.Next
				fast = fast.Next
			}
			return head
		}
	}
	return nil
}
