package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeatNumberInArray(t *testing.T) {
	arr := []int{2, 3, 1, 0, 2, 5}
	arr1 := []int{3, 1, 0, 2, 2, 5}
	arr2 := []int{3, 1, 0, 2, 4}
	assert.True(t, repeatNumberInArray(arr) == 2)
	assert.True(t, repeatNumberInArray(arr1) == 2)
	assert.True(t, repeatNumberInArray(arr2) == -1)
}
