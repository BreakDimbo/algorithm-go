package linkedList

/*
	solution 1: 遍历，主要是进位问题，使用 carry 变量保存每次的进位
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur, p, q := dummy, l1, l2
	carry := 0

	for p != nil || q != nil {
		x, y := 0, 0
		if p != nil {
			x = p.Val
			p = p.Next
		}

		if q != nil {
			y = q.Val
			q = q.Next
		}

		sum := x + y + carry
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	// 注意最后一个节点的进位问题
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
