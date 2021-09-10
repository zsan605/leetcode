package problems

func strStr(haystack string, needle string) int {
	nLen := len(needle)
	hLen := len(haystack)
	if nLen == 0 {
		return 0
	}

	if hLen == 0 || nLen > hLen {
		return -1
	}

	var isFlag = true
	for i := 0; i < hLen && nLen+i <= hLen; i++ {
		var j = 0
		isFlag = true
		for j = 0; j < nLen; j++ {
			if i+j < len(haystack) && haystack[i+j] != needle[j] {
				isFlag = false
				break
			}
		}
		if j == nLen && isFlag {
			return i
		}
	}
	return -1
}

func strStr1(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

func strStr2(haystack, needle string) int {
	hLen, nLen := len(haystack), len(needle)

	if nLen == 0 {
		return 0
	}

	if hLen==0 && nLen !=0 {
		return -1
	}
	var pi = make([]int, nLen)
	for i, j := 1, 0; i < nLen; i++ {
		if j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}

		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}

	for i, j := 0, 0; i < hLen; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}

		if haystack[i] == needle[j] {
			j++
		}

		if j == nLen {
			return i - nLen + 1
		}
	}
	return -1
}
