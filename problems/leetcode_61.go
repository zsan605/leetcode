package problems

func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var num = 1
	var iter = head
	for iter.Next != nil {
		num++
		iter = iter.Next
	}

	n := num - k % num
	if n == num {
		return head
	}

	iter.Next = head
	for n >0 {
		iter = iter.Next
		n--
	}
	ret := iter.Next
	iter.Next = nil
	return ret
}
