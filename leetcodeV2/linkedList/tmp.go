package linkedList

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre, cur *ListNode
	pre, cur = nil, head

	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head

	for cur != nil && cur.Next != nil {
		cur.Next, cur.Next.Next, pre.Next, cur, pre = cur.Next.Next, cur, cur.Next, cur.Next.Next, cur
	}
	return dummy.Next
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}
			return head
		}
	}
	return nil
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	start := dummy

	for {
		p := start
		start = p.Next
		tail, mv := p, p

		for i := 0; i < k && tail != nil; i++ {
			tail = tail.Next
		}

		if tail == nil {
			break
		}

		for i := 0; i < k-1; i++ {
			mv = p.Next
			p.Next = mv.Next
			mv.Next = tail.Next
			tail.Next = mv
		}
	}
	return dummy.Next
}
