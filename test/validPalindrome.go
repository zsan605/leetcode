package test

func validPalindrome(s string) bool {

	i, j, isPalindrome := validPalindromeFunc(s)
	if isPalindrome {
		return true
	}

	_, _, isPalindrome = validPalindromeFunc(s[i+1:j+1])
	if isPalindrome {
		return true
	}

	_, _, isPalindrome = validPalindromeFunc(s[i:j])
	if isPalindrome {
		return true
	}

	return false
}

func validPalindromeFunc(temp string) (int, int, bool) {
	i, j := 0, len(temp) -1
	for i < j {
		if temp[i] != temp[j] {
			return i, j, false
		}
		i++
		j--
	}
	return i, j, true
}