package problems

/*
i=[1, n]
fmax(i) = max{fmax(i-1)*ai, fmin(i-1)*ai, ai}
fmin(i) = min{fmax(i-1)*ai, fmin(i-1)*ai, ai}
*/
func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}
