package problems
/*
206.反转链表
思路1：头插法
*/

func reverseList_1(head *ListNode) *ListNode {
	var prev *ListNode
	var curr = head
	for curr!= nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func reverseList(head *ListNode) *ListNode {

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
