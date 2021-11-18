package problems

import "testing"

func TestCombinationSum2(t *testing.T) {
	temp := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	ret := combinationSum2(temp, target)
	t.Log("xxxxxxxxx")
	t.Log(ret)
}
