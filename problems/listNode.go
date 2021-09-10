package problems

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(data []int) *ListNode {
	var lastNode *ListNode
	for i := len(data) - 1; i >= 0; i-- {
		node := &ListNode{
			Val:  data[i],
			Next: lastNode,
		}
		lastNode = node
	}
	return lastNode
}

func PrintList(head *ListNode){
	for head != nil {
		fmt.Print(head.Val, "->")
		head = head.Next
	}
	fmt.Println("")
}