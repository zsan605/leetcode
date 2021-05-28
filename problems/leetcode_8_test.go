package problems

import "testing"

func TestMyAtoi(t *testing.T) {
	str := "42"
	result := myAtoi(str)
	t.Log(result)
}
