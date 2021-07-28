package problems

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	for i, num := range nums {
		if num >= target {
			return i
		}
	}

	return len(nums)
}

// func searchInsert1(nums []int, target int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	mid := len(nums) / 2

// 	for i, num := range nums {
// 		if num >= target {
// 			return i
// 		}
// 	}

// 	return len(nums)
// }
