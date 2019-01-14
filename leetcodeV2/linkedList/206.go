package linkedList

type ListNode struct {
	Next *ListNode
	val  int
}

func reverseList(head *ListNode) *ListNode {
	var cur, prev *ListNode = head, nil
	for cur != nil {
		cur, cur.Next, prev = cur.Next, prev, cur
	}
	return prev
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
