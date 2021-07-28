package problems

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(strs)
	t.Log(result)
}
