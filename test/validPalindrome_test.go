package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	assert.True(t, validPalindrome("abcddba"))
}
