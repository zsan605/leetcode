package problems

import "testing"

func TestTrap(t*testing.T){
	//var arr = []int{0,1,0,2,1,0,1,3,2,1,2,1}
	var arr = []int{4, 2, 3}
	ret :=trap(arr)
	t.Log(ret)
}
