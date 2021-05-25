package leetcode

import (
	"math"
)

/*
给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a*a + b*b = c 。

示例 1：

输入：c = 5
输出：true
解释：1 * 1 + 2 * 2 = 5
示例 2：

输入：c = 3
输出：false
示例 3：

输入：c = 4
输出：true
示例 4：

输入：c = 2
输出：true
示例 5：

输入：c = 1
输出：true

提示：
0 <= c <= 231 - 1

*/

func judgeSquareSum(c int) bool {

	i, j := 0, int(math.Sqrt(float64(c)))

	target := 0
	for i <= j {
		target = i*i + j*j
		if c == target {
			return true
		} else if target > c {
			j--
		} else {
			i++
		}
	}
	return false
}
