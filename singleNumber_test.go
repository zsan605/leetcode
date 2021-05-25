package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	var nums = []int{4, 1, 2, 1, 2}
	assert.True(t, singleNumber(nums) == 4)
}
