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

func reverseList(head *ListNode) *ListNode {
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
		pre.Next, cur.Next, cur.Next.Next = cur.Next, cur.Next.Next, cur
		pre, cur = cur, cur.Next
	}
	return dummy.Next
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// x1+x2+x3+x2 = 2(x1+x2) => x3 = x1
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
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

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	start := dummy
	var p, mv, tail *ListNode

	for {
		p = start
		mv, tail = p, p
		start = p.Next

		// find the tail
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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return divideAndConq(lists, 0, len(lists)-1)
}

func divideAndConq(lists []*ListNode, l, r int) *ListNode {
	if l >= r {
		return lists[l]
	}
	mid := l + ((r - l) >> 1)
	left := divideAndConq(lists, l, mid)
	right := divideAndConq(lists, mid+1, r)

	return mergeListx(left, right)
}

func mergeListx(l, r *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l != nil && r != nil {
		if l.Val < r.Val {
			cur.Next = l
			l = l.Next
		} else {
			cur.Next = r
			r = r.Next
		}
		cur = cur.Next
	}
	if l != nil {
		cur.Next = l
	}
	if r != nil {
		cur.Next = r
	}
	return dummy.Next
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		pre, head, head.Next = head, head.Next, pre
	}
	return pre
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head
	for cur != nil && cur.Next != nil {
		pre.Next, cur.Next, cur.Next.Next = cur.Next, cur.Next.Next, cur
		pre, cur = cur, cur.Next
	}
	return dummy.Next
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 2(x1 + x2) = x1 + x2 + x3 + x2 -> x1 = x3
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			for fast != head {
				fast = fast.Next
				head = head.Next
			}
			return fast
		}
	}
	return nil
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	anchor := dummy
	var preHead, mover, tail *ListNode

	for {
		preHead = anchor
		anchor = preHead.Next
		tail = preHead

		for i := 0; i < k && tail != nil; i++ {
			tail = tail.Next
		}

		if tail == nil {
			break
		}

		for i := 0; i < k-1; i++ {
			mover = preHead.Next
			preHead.Next = mover.Next
			mover.Next = tail.Next
			tail.Next = mover
		}
	}
	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return divideNConq(lists, 0, len(lists)-1)
}

func divideNConq(lists []*ListNode, lo, hi int) *ListNode {
	if lo == hi {
		return lists[lo]
	}
	mid := lo + ((hi - lo) >> 1)
	left := divideNConq(lists, lo, mid)
	right := divideNConq(lists, mid+1, hi)

	return mergex(left, right)
}

func mergex(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	list := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			list.Next = left
			left = left.Next
		} else {
			list.Next = right
			right = right.Next
		}
		list = list.Next
	}
	if left != nil {
		list.Next = left
	} else {
		list.Next = right
	}
	return dummy.Next
}
