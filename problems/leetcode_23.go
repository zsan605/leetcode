package problems

/*
前提：合并两个有序链表
方法：归并
*/
func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	// var head *models.ListNode
	// for _, list := range lists {
	// 	head = mergeTwoLists(head, list)
	// }

	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}

	if l > r {
		return nil
	}

	mid := (l + r) >> 1

	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}
