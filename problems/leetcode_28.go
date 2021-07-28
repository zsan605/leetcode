package problems

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) == 0 || len(needle) > len(haystack) {
		return -1
	}

	var isFlag = false
	for i := 0; i < len(haystack); i++ {
		var j = 0
		for j = 0; j < len(needle); j++ {
			if i+j < len(haystack) && haystack[i+j] != needle[j] {
				isFlag = false
				break
			}
		}
		if j == len(needle) && isFlag {
			return i
		}
	}
	return -1
}
