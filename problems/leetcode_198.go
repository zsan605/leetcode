package problems

func rob(nums []int) int {

	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	dp_2, dp_1 := nums[0], max(nums[0], nums[1])
	// var ans = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp_i := max(dp_2+nums[i], dp_1)
		dp_2, dp_1 = dp_1, dp_i
		// ans = max(ans, dp_i)
	}
	return dp_1
}
