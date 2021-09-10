package problems

import "testing"

func TestReverseWords(t *testing.T) {
	//str := "the sky is blue"
	str := "a good   example"
	ret := reverseWords(str)
	t.Log(ret)

}
