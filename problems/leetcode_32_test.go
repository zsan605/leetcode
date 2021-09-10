package problems

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	var str = "()(()"
	t.Log(longestValidParentheses(str))
}
