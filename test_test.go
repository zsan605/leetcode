package leetcode

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestTestA(t *testing.T) {
// 	assert.True(t, testA())
// }

func TestWordIndex(t *testing.T) {
	// assert.True(t, wordIndex(64) == 1)

	ret, _ := regexp.MatchString(`\d{11}$`, "00000000000")
	assert.True(t, ret == true)
	// t.Log(wordIndex(6))
	// fmt.Println(wordIndex(6))
}
