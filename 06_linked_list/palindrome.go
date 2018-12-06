package linkedlist

func palindrome(s string) bool {
	list := NewSingleLinkedList()
	for _, c := range s {
		list.InsertValueToTail(string(c))
	}

	if list.length == 1 {
		return true
	}

	if list.length == 2 {
		return list.head == list.head.next
	}

	// 如果是字符数偶数，在 j 遍历到最后一个元素时停止
	// 1，2|，3，4
	if list.length%2 == 0 {
		i := list.head
		j := list.head.next
		for j.next != nil {

			// i 跳一步，j 跳两步
			i = i.next
			j = j.next.next
		}
	}

	// 如果字符数是奇数

}
