package problems

/*
	双指针p1,p2
	每次移动两个节点
*/
func swapPairs(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	p1, p2 := head, head.Next
	for p1 != nil && p2 != nil {
		p1.Val, p2.Val = p2.Val, p1.Val

		if p2.Next == nil {
			break
		}

		p1 = p2.Next
		p2 = p1.Next
	}

	return head
}
