package problems

/*
地址：https://leetcode-cn.com/problems/squares-of-a-sorted-array/
思路：双指针
1. 首尾指针向中间靠拢
2. 计算首尾指针位置的数的乘积
3. 如果尾指针处的大，则尾指针前移
4. 如果首指针处的大，则首指针后移
5. 直至首尾指针相等
*/

func sortedSquares(nums []int) []int {

	i, j, p := 0, len(nums)-1, len(nums)-1
	var ret = make([]int, len(nums))
	for i <= j {
		if nums[j]*nums[j] > nums[i]*nums[i] {
			ret[p] = nums[j] * nums[j]
			j--
		} else{
			ret[p] = nums[i] * nums[i]
			i++
		}
		p--
	}
	return ret
}

