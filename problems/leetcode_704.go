package problems

/*
地址：https://leetcode-cn.com/problems/binary-search/
思路：
由于nums是升序数组
1. 找到nums的中间位置
2. 判断中间数和目标数的大小
3. 如果目标数等于中间数据，返回中间位置
4. 如果目标数大于中间数，则在中间数到nums结尾的区间进行二分查找
5. 如果目标数小于中间数，则在nums开始到中间数的区间进行二分查找
6. 如果没有找到返回-1

边界：
1. 数组判空
2. 数组边界
3. nums 为奇，偶时
*/
func search(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return -1
	}
	return searchHelper(nums, 0, numsLen-1, target)
}

func searchHelper(nums []int, start, end, target int) int {

	mid := (end + start) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target && mid+1 <= end {
		return searchHelper(nums, mid+1, end, target)
	} else if nums[mid] > target && mid-1 >= start {
		return searchHelper(nums, start, mid, target)
	}
	return -1
}

/*
非递归思路，双指针
1. 设置双指针，指向头尾
2. 判断尾指针是否在首指针之后
3. 如果不是，返回-1
4. 如果是，计算中间指针
5. 如果中间指针等于目标值，返回中间指针位置
6. 如果中间指针小于目标值，设置首指针为中间位置+1
7. 否知设置尾指针为中间位置-1
*/
func search2(nums []int, target int) int {

	start, end := 0, len(nums)-1

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
