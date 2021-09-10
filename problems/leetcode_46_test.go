package problems

import "testing"

func TestPermute(t *testing.T){
	var nums = []int{1,2,3}
	ret :=permute(nums)
	t.Log(ret)
}