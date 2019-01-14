package linkedList

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyNode := &ListNode{Next: head}
	start := dummyNode

	for {
		p := start
		c, n := p, p
		start = p.Next

		// find the k+1 node
		for i := 0; i < k && n != nil; i++ {
			n = n.Next
		}
		if n == nil {
			break
		}

		for i := 0; i < k-1; i++ {
			c = p.Next
			p.Next = c.Next
			c.Next = n.Next
			n.Next = c
		}
	}
	return dummyNode.Next
}
