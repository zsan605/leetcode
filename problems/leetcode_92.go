package problems

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if head == nil {
		return nil
	}

	var stack []int
	var i, num = 0, 0

	ptr1, ptr2 := head, head

	for ptr1 != nil {
		num++

		if num < left {
			ptr2 = ptr2.Next
		}

		if num >= left && num <= right {
			stack = append(stack, ptr1.Val)
			i++
		}
		ptr1 = ptr1.Next
	}
	for i >= 1 {
		ptr2.Val = stack[i-1]
		ptr2 = ptr2.Next
		i--
	}
	return head
}

func reverseBetween1(head *ListNode, left int, right int) *ListNode {

	if head == nil {
		return nil
	}

	ptr1, ptr2 := head, head

	i := 1
	for i != left {
		ptr1 = ptr1.Next
		i++
	}

	for i != right {
		ptr2 = ptr2.Next
		i++
	}

	j := 1
	for left+j < right {
		ptr1.Next = ptr2.Next
		ptr2.Next = ptr1
	}

	// for i >= 1 {
	// 	ptr2.Val = stack[i-1]
	// 	ptr2 = ptr2.Next
	// 	i--
	// }
	return head
}
