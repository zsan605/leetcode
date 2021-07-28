package problems

import (
	"leetcode/models"
	"testing"
)

func TestReverseKGroup(t *testing.T) {

	head := models.CreateList([]int{1,2,3,4,5})
	models.PrintList(head)
	ret :=reverseKGroup(head, 3)
	models.PrintList(ret)

}
