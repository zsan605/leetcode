package problems

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	str := "au"

	temp := lengthOfLongestSubstring(str)
	t.Log(temp)
}
