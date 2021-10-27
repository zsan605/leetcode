package test

/*
给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。

函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。

说明:

返回的下标值（index1 和 index2）不是从零开始的。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
示例:

输入: numbers = [2, 7, 11, 15], target = 9
输出: [1,2]
解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。

*/

func twoSum0(numbers []int, target int) []int {

	if len(numbers) < 2 {
		return nil
	}

	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i, j}
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}
	return nil
}

func twoSum(nums []int, target int) []int {

	l := len(nums)

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {

	numsMap := make(map[int]int)
	for i, v := range nums {
		numsMap[v] = i
	}

	for i := 0; i < len(nums); i++ {
		num := target - nums[i]
		j, ok := numsMap[num]
		if ok && j != i {
			return []int{i, j}
		}
	}

	return []int{}
}

func twoSum3(nums []int, target int) []int {

	numsMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		num := target - nums[i]
		j, ok := numsMap[num]
		if ok {
			return []int{i, j}
		}
		numsMap[nums[i]] = i
	}

	return []int{}
}

func twoSum4(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	var temp int
	for i < j {
		temp = nums[i] + nums[j]
		if temp == target {
			i++
			j++
			return []int{i, j}
		} else if temp < target {
			i++
		} else {
			j--
		}
	}
	return []int{}
}
