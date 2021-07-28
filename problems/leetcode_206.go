package problems

import "leetcode/models"

func reverseList(head *models.ListNode) *models.ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	p1, p2 := head, head.Next
	head.Next = nil

	for p2 != nil {
		p3 := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}
	return p1
}
