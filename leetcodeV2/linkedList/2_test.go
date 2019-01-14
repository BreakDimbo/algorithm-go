package linkedList

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 5}
	l1.Next = &ListNode{Val: 6}
	l1.Next.Next = &ListNode{Val: 4}

	l2 := &ListNode{Val: 2}
	l2.Next = &ListNode{Val: 4}
	l2.Next.Next = &ListNode{Val: 3}

	addTwoNumbers(l1, l2)
}
