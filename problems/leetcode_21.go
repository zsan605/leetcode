package problems

import "leetcode/models"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
 解题思路：双指针
 遍历一个链表，向另一个链表中插入

*/
func mergeTwoLists(l1 *models.ListNode, l2 *models.ListNode) *models.ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var hL, pL, tL *models.ListNode

	if l1.Val <= l2.Val {
		hL = l1
		pL = l1
		tL = l2
	} else {
		hL = l2
		pL = l2
		tL = l1
	}

	for tL != nil {
		if pL.Next == nil {
			pL.Next = tL
			break
		}

		if tL.Val >= pL.Val && tL.Val >= pL.Next.Val {
			pL = pL.Next
			continue
		}
		tLN := tL.Next
		tL.Next = pL.Next
		pL.Next = tL
		tL = tLN
	}
	return hL
}
