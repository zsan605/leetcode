package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasCycle(t *testing.T) {
	var node1 *ListNode
	//node7 := &ListNode{
	//	Val: 7,
	//	Next: node1,
	//}
	//node6 := &ListNode{
	//	Val: 6,
	//	Next: node7,
	//}
	//node5 := &ListNode{
	//	Val: 5,
	//	Next: node6,
	//}
	//node4 := &ListNode{
	//	Val: 4,
	//	Next: node5,
	//}
	//node3 := &ListNode{
	//	Val: 3,
	//	Next: node4,
	//}
	//node2 := &ListNode{
	//	Val: 2,
	//	Next: nil,
	//}
	node1 = &ListNode{
		Val:  1,
		Next: nil,
	}

	head := &ListNode{
		Val: 1,
		Next: node1,
	}
	assert.True(t, hasCycle(head))
}