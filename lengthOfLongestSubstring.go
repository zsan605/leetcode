package leetcode

// 最长子串
// 滑动窗口
// 1. i, j分别为滑动窗口的左右边界
// 2. index 为滑动窗口的游标，从i到j
// 3. 判断index对应的值和滑动窗口后面的值是否相等
// 4. 如果相等，将滑动窗口的左边界设为index+1，右边界设为j+1
// 5. 如果不相等，index+1,直至index等于j,然后有边界设为j+1
func lengthOfLongestSubstring(s string) int {
	i, l := 0, len(s)

	if l == 0 || l == 1 {
		return l
	}

	maxLength := 0
	for j := 1; j <= l; j++ {

		length := j - i
		if length > maxLength {
			maxLength = length
		}

		if j == l {
			break
		}

		for index := i; index < j; index++ {
			if s[index : index+1] == s[j : j+1] {
				i = index + 1
				break
			}
		}
	}

	return maxLength
}
