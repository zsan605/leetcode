package problems

import "testing"

func TestConvert(t *testing.T) {
	str := "PAYPALISHIRING"
	str1 := convert(str, 3)
	t.Log(str1)
}
