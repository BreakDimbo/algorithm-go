package linkedList

/*
	solution 1: brute force
	solution 2: hash
	solution 3: two pointers, redirect A to B when A get end, redirect B to A head when B get end
*/

/*
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var lengthA, lengthB int

	nodeA, nodeB := headA, headB
	for nodeA != nil || nodeB != nil {
		if nodeA != nil {
			nodeA = nodeA.Next
			lengthA++
		}

		if nodeB != nil {
			nodeB = nodeB.Next
			lengthB++
		}
	}

	nodeA, nodeB = headA, headB
	if lengthA > lengthB {
		for i := 0; i < lengthA-lengthB; i++ {
			nodeA = nodeA.Next
		}
	} else if lengthA < lengthB {
		for i := 0; i < lengthB-lengthA; i++ {
			nodeB = nodeB.Next
		}
	}
	for nodeA != nil && nodeB != nil {
		if nodeA == nodeB {
			return nodeA
		}
		nodeA = nodeA.Next
		nodeB = nodeB.Next
	}
	return nil
}
*/

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
