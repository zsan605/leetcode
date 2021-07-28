package problems

func longestCommonPrefix(strs []string) string {

	strsLen := len(strs)
	if strsLen == 0 {
		return ""
	} else if strsLen == 1 {
		return strs[0]
	}

	mid := strsLen / 2
	left := longestCommonPrefix(strs[0:mid])
	right := longestCommonPrefix(strs[mid:])

	var i int
	for i = 0; i < len(left) && i < len(right); i++ {
		if left[i] != right[i] {
			return left[0:i]
		}
	}
	return left[0:i]
}
