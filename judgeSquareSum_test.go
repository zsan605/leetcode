package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJudgeSquareSum(t *testing.T) {

	assert.True(t, judgeSquareSum(5))
	assert.True(t, judgeSquareSum(4))
	assert.False(t, judgeSquareSum(3))
	assert.True(t, judgeSquareSum(2))
	assert.True(t, judgeSquareSum(1))
	assert.False(t, judgeSquareSum(2147482647))

}