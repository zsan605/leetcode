package problems

// 暴力 O(n^2)
func twoSum(nums []int, target int) []int {
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 以空间换时间
func twoSum2(nums []int, target int) []int {

	if len(nums) < 2 {
		return nil
	}

	var numsMap = make(map[int]int)
	for i, val := range nums {
		numsMap[target-val] = i
	}

	for i, val := range nums {
		j, ok := numsMap[val]
		if ok && i != j {
			return []int{i, j}
		}
	}
	return nil
}

// 以空间换时间，优化版
func twoSum3(nums []int, target int) []int {
	tempMap := make(map[int]int)

	for i, val := range nums {
		p, ok := tempMap[target-val]
		if ok {
			return []int{i, p}
		}
		tempMap[val] = i
	}
	return nil
}
