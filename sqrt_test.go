package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrt(t *testing.T) {
	assert.True(t, sqrt(8) == 2)
	assert.True(t, sqrt(4) == 2)
	assert.True(t, sqrt(10) == 3)
	assert.True(t, sqrt(1) == 1)
	assert.True(t, sqrt(2) == 1)
	assert.True(t, sqrt(100) == 10)
}

func TestA(t *testing.T) {
	var aa = 1290 >> 6
	assert.True(t, aa == 0)
}
