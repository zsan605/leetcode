package problems

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	p, q := head, head.Next

	for p != q {
		if q == nil || q.Next == nil {
			return false
		}
		p = p.Next
		q = q.Next.Next
	}

	return true
}
