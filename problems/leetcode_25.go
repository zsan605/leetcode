package problems

func reverseKGroup(head *ListNode, k int) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	p1 := head
	i := 1
	for i = 1; i < k && p1 != nil; i++ {
		p1 = p1.Next
	}

	if k == i && p1 != nil {
		p2 := p1.Next
		p1.Next = nil
		temp := reverseList(head)
		head.Next = reverseKGroup(p2, k)
		return temp
	} else {
		return head
	}

}
