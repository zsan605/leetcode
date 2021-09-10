package problems

func deleteDuplicates_2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head = &ListNode{
		Val:  0,
		Next: head,
	}
	iter := head

	for iter != nil && iter.Next != nil && iter.Next.Next != nil {

		if iter.Next.Val == iter.Next.Next.Val {
			cur := iter.Next.Next
			for cur != nil {
				if cur.Val != iter.Next.Val {
					break
				}
				cur = cur.Next
			}
			iter.Next = cur
		} else {
			iter = iter.Next
		}
	}

	return head.Next
}
