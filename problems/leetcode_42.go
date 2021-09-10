package problems

func trap(height []int) int {

	if len(height) <= 2 {
		return 0
	}

	var area int
	left, right, leftMax, rightMax := 0, len(height)-1, 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			area = area + leftMax - height[left]
			left++
		} else {
			area = area + rightMax - height[right]
			right--
		}
	}

	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
