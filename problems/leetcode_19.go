package problems

/*
双指针解法
p1 = head
p2 = head+n
当p2到链表结尾，p1刚好到到时第n个
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	p1, p2 := head, head

	for i := 1; i <= n; i++ {
		if p2 == nil {
			return nil
		}
		p2 = p2.Next
	}

	if p2 == nil {
		return head.Next
	}

	for p2.Next != nil {
		p2 = p2.Next
		p1 = p1.Next
	}

	p1.Next = p1.Next.Next

	return head
}
