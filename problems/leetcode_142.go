package problems

func detectCycle(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	}

	slow, faster := head, head.Next

	for slow != faster {
		if faster == nil || faster.Next == nil {
			return nil
		}
		slow = slow.Next
		faster = faster.Next.Next
	}

	ptr := head

	for ptr != slow {
		ptr = ptr.Next
		slow = slow.Next
	}
	return ptr
}
