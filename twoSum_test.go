package leetcode

import (
	"fmt"
	"testing"
)


func TestTwoSum(t *testing.T) {
	nums := []int{2,7,11,15}
	target := 9

	fmt.Println(twoSum0(nums, target))
	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSum2(nums, target))
	fmt.Println(twoSum3(nums, target))
	fmt.Println(twoSum4(nums, target))
}
