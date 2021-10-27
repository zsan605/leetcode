package test

import (
	"fmt"
	"strings"
)

/*
编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
示例 1：

输入："hello"
输出："holle"
示例 2：

输入："leetcode"
输出："leotcede"


提示：

元音字母不包含字母 "y" 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-vowels-of-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reverseVowels(s string) string {
	fmt.Println([]byte("aeiouAEIOU"))
	vowelMap := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}
	i, j := 0, len(s)-1

	for i < j {
		a, b := strings.ToLower(s[i:i+1]), strings.ToLower(s[j:j+1])
		if vowelMap[a] && vowelMap[b] {
			s = s[0:i] + s[j:j+1] + s[i+1:j] + s[i:i+1] + s[j+1:]
			i++
			j--
			//s[i:i+1], s[j:j+1] = s[i:i+1], s[j:j+1]
		}
		if !vowelMap[a] {
			i++
		}
		if !vowelMap[b] {
			j--
		}
	}
	return s
}

func reverseVowels1(s string) string {

	vowelMap := map[uint8]bool{
		byte('a'): true,
		byte('e'): true,
		byte('i'): true,
		byte('o'): true,
		byte('u'): true,
		byte('A'): true,
		byte('E'): true,
		byte('I'): true,
		byte('O'): true,
		byte('U'): true,
	}
	i, j := 0, len(s)-1
	temp := []byte(s)
	for i < j {
		if vowelMap[temp[i]] && vowelMap[temp[j]] {
			temp[i], temp[j] = temp[j], temp[i]
			i++
			j--
		}
		if !vowelMap[temp[i]] {
			i++
		}
		if !vowelMap[temp[j]] {
			j--
		}
	}
	return string(temp)
}

func isUnique(astr string) bool {
	if len(astr) == 0 || len(astr) == 1 {
		return true
	}

	for i := 0; i < len(astr)-2; i++ {
		for j := i + 1; j < len(astr)-1; j++ {
			if astr[i:i+1] == astr[j:j+1] {
				return false
			}
		}
	}
	return true
}
