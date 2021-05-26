package problems

import (
	"reflect"
	"unsafe"
)

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
示例 4:

输入: s = ""
输出: 0


提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
解题思路：
	滑动窗口
	0. 初始化子串最大长度为0
	1. 用i，j标识子串的首尾指针
	2. 用k,标识待加入子串的元素（紧邻子串尾部）
	3. 判断待加入子串是否在子串中
	3.1. 如果不在
		3.1.1 子串向后拓展一位，j=k
	3.2. 如果存在
		3.2.1 将子串长度与最大子串长度比较，如果大于最大子串则替换
		3.2.2 将k所指字符替换当前子串，i=k, j=k+1
	4. k向后移一位,k=k+2
	5. 跳转到3执行
边界条件：
	字符串到最后
*/

func lengthOfLongestSubstring(s string) int {

	maxLen, l := 1, len(s)
	if l <= 1 {
		return l
	}

	sBytes := (*[]byte)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s))))

	i, j, k := 0, 1, 1
	for k < l {
		index := isInStr(sBytes, i, j, k)
		if index == -1 {
			k += 1
			j += 1
			continue
		} else {
			temp := j - i + 1
			if maxLen < temp {
				maxLen = temp
			}
			i = index + 1
		}
	}
	temp := j - i
	if maxLen < temp {
		maxLen = temp
	}
	return maxLen
}

func isInStr(s *[]byte, i, j, k int) int {
	for i != j {
		if (*s)[i] == (*s)[k] {
			return i
		}
		i += 1
	}
	return -1
}
