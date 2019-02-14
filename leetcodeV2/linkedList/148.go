package linkedList

/*
	solution 1: recersive
	solution 2: iterative
*/

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// split the list
	var prev, slow, fast *ListNode
	slow, fast = head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = nil
	left := sortList(head)
	right := sortList(slow)

	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	l := &ListNode{}
	p := l

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	}

	if l2 != nil {
		p.Next = l2
	}

	return l.Next
}
