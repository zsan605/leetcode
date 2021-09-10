package problems

import (
	"strings"
)

func reverseWords(s string) string {

	temp := strings.Split(strings.Trim(s, " "), " ")
	temp = reverseWordsArr(temp)

	return strings.Join(temp, " ")
}

func reverseWordsArr(arr []string) []string {
	var ret []string
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != "" {
			ret = append(ret, arr[i])
		}
	}

	return ret
}
