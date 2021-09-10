package problems

/*
https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	pA, pB := headA, headB

	for {

		if pA == nil && pB == nil {
			break
		} else if pA == nil && pB != nil {
			pA = headB
		} else if pA != nil && pB == nil {
			pB = headA
		}

		if pA == pB {
			return pA
		}

		pA = pA.Next
		pB = pB.Next
	}

	return nil
}
