package problems

/*
地址: https://leetcode-cn.com/problems/search-insert-position/
思路：二分查找
1.
*/
func searchInsert(nums []int, target int) int {

	left, right := 0, len(nums)-1

	var mid int

	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1 // target 在区间[mid+1, right]
		} else if nums[mid] > target {
			right = mid - 1 // target 在区间[left, mid-1]
		}
	}

	return left
}
