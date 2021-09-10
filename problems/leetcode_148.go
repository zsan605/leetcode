package problems

// 冒泡
//func sortList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	p1 := head
//	for p1.Next != nil {
//		p2 := p1.Next
//		for p2 != nil {
//			if p1.Val > p2.Val {
//				p1.Val, p2.Val = p2.Val, p1.Val
//			}
//			p2 = p2.Next
//		}
//		p1 = p1.Next
//	}
//	return head
//}

//
//func sortList(head *ListNode) *ListNode {
//
//	var arr []*ListNode
//	cur := head
//
//	if head == nil {
//		return nil
//	}
//
//	for cur != nil {
//		arr = append(arr, cur)
//		cur = cur.Next
//	}
//
//	arr = mergeSort(arr)
//	for i := 0; i < len(arr)-1; i++ {
//		arr[i].Next = arr[i+1]
//	}
//
//	return arr[0]
//}
//
//func mergeSortList(head, tail *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return nil
//	}
//
//	mid
//
//	return merge1(left, right)
//}
//
//func merge1(left []*ListNode, right []*ListNode) []*ListNode {
//
//	leftN, rightN := len(left), len(right)
//	l, r := 0, 0
//
//	var ret []*ListNode
//	for l < leftN && r < rightN {
//		if left[l].Val < right[r].Val {
//			ret[l+r] = left[l]
//			l++
//		} else {
//			ret[l+r] = right[r]
//			r++
//		}
//	}
//	ret = append(ret, left[l:]...)
//	ret = append(ret, right[r:]...)
//	return ret
//}
