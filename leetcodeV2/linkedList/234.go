package linkedList

/*
	solution 1: use array to store
	solution 2: reverse the linkedList
	solution 3: slow and fast + reverse
*/

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// odd nodes, move slow to make right smaller
	if fast != nil {
		slow = slow.Next
	}

	slow = reverse(slow)
	fast = head

	for slow != nil {
		if fast.Val != slow.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}

func reverse(n *ListNode) *ListNode {
	var pre, node *ListNode
	pre, node = nil, n
	for node != nil {
		pre, node.Next, node = node, pre, node.Next
	}
	return pre
}
