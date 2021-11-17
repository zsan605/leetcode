package problems

/*
75. 颜色分类
*/

/*
思路1. 双指针，大的放右端， 小的放左端
*/
func sortColors(nums []int)  {
	index, left, right := 0, 0, len(nums)

	for left < right {
		if nums[index] < 1 {
			nums[left], nums[index] = nums[index], nums[left]
			left++
		} else if nums[index] == 1 {
			index++
		} else {
			nums[right], nums[index] = nums[index], nums[right]
			right--
		}
	}
}