package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	data := map[string]int{
		"aaaaaaub":3,
		"au":2,
		"abcabcbb": 3,
		"bbb":    1,
		"pwwkew":   3,
	}

	for key, value := range data {
		assert.True(t, lengthOfLongestSubstring(key) == value)
	}

}
