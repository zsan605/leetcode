package problems

func canJump(nums []int) bool {
	rightMost, n := 0, len(nums)
	for i := 0; i < len(nums); i++ {
		if i <= rightMost {
			rightMost = max(rightMost, i+nums[i])
			if rightMost >= n-1 {
				return true
			}
		}
	}

	return false
}
