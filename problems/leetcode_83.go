package problems

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	iter := head

	for iter != nil && iter.Next != nil {

		if iter.Val == iter.Next.Val {
			iter.Next = iter.Next.Next
		} else {
			iter = iter.Next
		}
	}

	return head
}
