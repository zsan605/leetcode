package problems

/*
i = [1, n]
fmax(i) = max{fmax(i-1)+ai, ai}
*/
func maxSubArray(nums []int) int {
	maxF, ans := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxF = max(nums[i], maxF+nums[i])
		ans = max(maxF, ans)
	}
	return ans
}
