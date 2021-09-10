package problems

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	nums1 := []int{-1,0,3,5,9,12}
	nums2 := []int{}
	nums3 := []int{-1}

	assert.Equal(t, search(nums1, 9), 4)
	assert.Equal(t, search(nums1, 13), -1)
	assert.Equal(t, search(nums2, 13), -1)
	assert.Equal(t, search(nums3, -1), 0)
	assert.Equal(t, search(nums3, -12), -1)
}

func TestSearch2(t *testing.T) {
	nums1 := []int{-1,0,3,5,9,12}
	nums2 := []int{}
	nums3 := []int{-1}

	assert.Equal(t, search2(nums1, 9), 4)
	assert.Equal(t, search2(nums1, 13), -1)
	assert.Equal(t, search2(nums2, 13), -1)
	assert.Equal(t, search2(nums3, -1), 0)
	assert.Equal(t, search2(nums3, -12), -1)
}