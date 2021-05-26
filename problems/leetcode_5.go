package problems

import (
	"reflect"
	"unsafe"
)

/*
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
示例 3：

输入：s = "a"
输出："a"
示例 4：

输入：s = "ac"
输出："a"


提示：

1 <= s.length <= 1000
s 仅由数字和英文字母（大写和/或小写）组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
解题思路：

*/

func longestPalindrome(s string) string {

	l := len(s)
	if l <= 1 {
		return s
	}
	left, right := 0, 0

	sBytes := (*[]byte)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s))))
	for i := 0; i < l; i++ {
		for j := l - 1; j > i; j-- {
			if (*sBytes)[i] == (*sBytes)[j] {
				isPali := isPalindrome(sBytes, i, j)
				if isPali {
					if j-i > right-left {
						left, right = i, j
					}
					break
				}
			}
		}
	}
	return s[left : right+1]
}

func isPalindrome(s *[]byte, i, j int) bool {
	for i < j {
		if (*s)[i] != (*s)[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 中心扩散法
func longestPalindrome2(s string) string {

	l := len(s)
	if l <= 1 {
		return s
	}
	left, right := 0, 0

	sBytes := (*[]byte)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s))))

	for i := 0; i < l-1; i++ {
		left1, right1 := expendStr(sBytes, l, i, i)
		left2, right2 := expendStr(sBytes, l, i, i+1)

		if right1-left1 > right-left {
			left, right = left1, right1
		}

		if right2-left2 > right-left {
			left, right = left2, right2
		}
	}
	return s[left : right+1]
}

func expendStr(s *[]byte, l, i, j int) (int, int) {
	for i >= 0 && j < l {
		if (*s)[i] != (*s)[j] {
			break
		}
		i--
		j++
	}
	return i + 1, j - 1
}

// 动态规划法

// func longestPalindrome3(s string) string {

// 	l := len(s)
// 	if l <= 1 {
// 		return s
// 	}
// 	left, right := 0, 0
// 	var dp [][]int
// 	for i := 0; i < l; i++ {
// 		dp[i][i] = 1
// 	}

// 	for i := 0; i < l; i++ {
// 		for j := 0; j < i; j++ {
// 			if
// 		}
// 	}

// 	return s[left : right+1]
// }
