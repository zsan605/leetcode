package problems

func removeElements(head *ListNode, val int) *ListNode {

	if head == nil {
		return nil
	}

	pre, cur := head, head.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}

	if head.Val == val {
		head = head.Next
	}

	return head
}
