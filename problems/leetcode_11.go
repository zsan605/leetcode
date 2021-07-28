package problems

// 双指针思路
// 证明指针移动
func maxArea(height []int) int {

	maxArea, area := 0, 0
	if len(height) <= 1 {
		return 0
	}

	left, right := 0, len(height)-1

	for left < right {
		if height[left] < height[right] {
			area = (right - left) * height[left]
			left++
		} else {
			area = (right - left) * height[right]
			right--
		}
		if area > maxArea {
			maxArea = area
		}

	}

	return maxArea
}
