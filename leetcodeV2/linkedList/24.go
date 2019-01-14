package linkedList

func swapPairs(head *ListNode) *ListNode {
	self := &ListNode{Next: head}
	prev := self
	for prev.Next != nil && prev.Next.Next != nil {
		a := prev.Next
		b := a.Next
		prev.Next, a.Next, b.Next = a.Next, b.Next, a
		prev = a
	}
	return self.Next
}

func swapPairsV2(head *ListNode) *ListNode {
	self := &ListNode{Next: head}
	prev, cur := self, head
	for cur != nil && cur.Next != nil {
		prev.Next, cur.Next, cur.Next.Next = cur.Next, cur.Next.Next, cur
		prev, cur = cur, cur.Next
	}
	return self.Next
}
