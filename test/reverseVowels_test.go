package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseVowels(t *testing.T) {
	assert.True(t, reverseVowels("hello") == "holle")
	assert.True(t, reverseVowels("aA") == "Aa")
	assert.True(t, reverseVowels("leetcode") == "leotcede")
}
