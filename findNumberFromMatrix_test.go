package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNumberFromMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	assert.True(t, findNumberFromMatrix(12, matrix))
	assert.False(t, findNumberFromMatrix(20, matrix))
	assert.False(t, findNumberFromMatrix(0, matrix))
}
