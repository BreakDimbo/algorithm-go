package linkedList

/*
	solution 1: reverse the list and find the nth from the head
	solution 2:
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	r := reverse(head)

	if n == 1 {
		tmp := r.Next
		r = nil
		return reverse(tmp)
	}

	var prev, cur *ListNode = nil, r
	for i := 0; i < n-1; i++ {
		prev = cur
		cur = cur.Next
	}
	prev.Next = cur.Next

	return reverse(r)
}

func reverse(node *ListNode) *ListNode {
	var prev, cur *ListNode = nil, node
	for cur != nil {
		prev, cur.Next, cur = cur, prev, cur.Next
	}
	return prev
}
