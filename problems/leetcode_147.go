package problems

func reorderList(head *ListNode) {

	if head == nil || head.Next == nil {
		return
	}

	slow, faster := head, head
	for faster.Next != nil && faster.Next.Next != nil {
		slow = slow.Next
		faster = faster.Next.Next
	}

	l1 := head
	l2 := slow.Next
	slow.Next = nil
	slowRev := reverseList1(l2)

	mergeList(l1, slowRev)
}

func reverseList1(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

func mergeList(l1, l2 *ListNode)  {
	var l1p, l2p *ListNode
	for l1 != nil && l2 != nil {
		l1p = l1.Next
		l2p = l2.Next

		l1.Next = l2
		l1 = l1p

		l2.Next = l1
		l2 = l2p
	}
}