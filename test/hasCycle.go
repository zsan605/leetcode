package test

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/

type ListNode struct {
     Val int
     Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	first, second := head, head.Next.Next
	for first != nil && second != nil {
		if first == second {
			return true
		}else {
			first = first.Next
			second = second.Next
			if second == nil {
				return false
			} else {
				second = second.Next
			}
		}
	}
	return false
}