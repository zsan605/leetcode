package test

func longestPalindrome(s string) string {

	l := len(s)
	if l == 0 || l == 1 {
		return s
	}

	max := 1
	h,g :=0, 1
	for i := 1; i < l-1; i++ {
		j, k := i-1, i+1
		if s[i] == s[k] {
			k++
		}
		for j > 0 && k < l {
			if s[j] != s[k] {
				j++
				k--
				break
			}
			k++
			j--
		}
		if max < k-j+1 {
			max = k-j+1
			h, g = j, k
		}
	}

	return s[h:g+1]
}
