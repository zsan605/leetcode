package problems

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例1：
2-->4-->3
5-->6-->4
7-->0-->8

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
解题思路：
	双指针遍历
情况分析：
	1. l1 和 l2 一样长
	2. l1 比 l2 短
	3. l2 比 l1 短
	4. 结果和最长的一样长
	5. 结果比最长的长
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var tempNode = new(ListNode)
	var head = tempNode
	var decade, unit int
	for l1 != nil || l2 != nil {
		var val int
		if l1 != nil && l2 != nil {
			val = l1.Val + l2.Val + decade
			l1, l2 = l1.Next, l2.Next
		} else if l1 == nil && l2 != nil {
			val = l2.Val + decade
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			val = l1.Val + decade
			l1 = l1.Next
		}

		decade, unit = val/10, val%10
		tempNode.Val = unit

		if l1 != nil || l2 != nil {
			tempNode.Next = new(ListNode)
			tempNode = tempNode.Next
		}
	}

	if decade != 0 {
		tempNode.Next = new(ListNode)
		tempNode.Next.Val = decade
	}
	return head
}
