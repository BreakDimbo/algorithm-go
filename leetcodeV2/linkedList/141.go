package linkedList

/*
	solution1: 使用快慢指针，看最终是否会相遇
	solution2: 使用 map 记录走过的 node 的
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil && slow != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func hasCycleV2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	visited := make(map[*ListNode]bool)
	for head != nil {
		if visited[head] {
			return true
		}
		visited[head] = true
		head = head.Next
	}
	return false
}
