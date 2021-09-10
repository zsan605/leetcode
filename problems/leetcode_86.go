package problems

func partition1(head *ListNode, x int) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	}

	h, pPre, p := head, head, head.Next

	for p != nil {
		if p.Val >= x {
			pPre = p
			p = p.Next
			continue
		}

		pPre.Next = p.Next

		if h.Val >= x {
			p.Next = h
			h = p
			head = h
			continue
		}

		p.Next = h.Next
		h.Next = p

		p = pPre.Next
	}

	return head
}

func partition(head *ListNode, x int) *ListNode {

	var small, big, smallP, bigP *ListNode

	for head != nil {
		if head.Val >= x {
			if big == nil {
				big = head
				bigP = head
			} else {
				bigP.Next = head
				bigP = bigP.Next
			}
		} else {
			if small == nil {
				small = head
				smallP = head
			} else {
				smallP.Next = head
				smallP = smallP.Next
			}
		}
		head = head.Next
	}

	if smallP != nil {
		smallP.Next = nil
	}
	if bigP != nil {
		bigP.Next = nil
	}

	if small == nil {
		return big
	}

	smallP.Next = big
	return small
}
