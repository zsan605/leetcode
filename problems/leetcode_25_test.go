package problems

import (
	"testing"
)

func TestReverseKGroup(t *testing.T) {

	head := CreateList([]int{1,2,3,4,5})
	PrintList(head)
	ret :=reverseKGroup(head, 3)
	PrintList(ret)

}
