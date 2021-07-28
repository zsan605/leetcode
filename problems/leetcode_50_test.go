package problems

import "testing"

func TestMyPow(t *testing.T) {
	t.Log(myPow(2, 10))
	t.Log(myPow(2.1, 3))
	t.Log(myPow(2.0, -2))

}