package test

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	str := "babad"
	str1 := "a"
	str2 := "ab"
	str3 := "aba"
	str4 := "abafadfasdf"

	fmt.Println(longestPalindrome(str))
	fmt.Println(longestPalindrome(str1))
	fmt.Println(longestPalindrome(str2))
	fmt.Println(longestPalindrome(str3))
	fmt.Println(longestPalindrome(str4))
}
