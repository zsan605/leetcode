package leetcode

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return getArrayMiddle(mergeArray(nums1, nums2))
}

func mergeArray(nums1 []int, nums2 []int) []int {

	if len(nums1) == 0 {
		return nums2
	}

	if len(nums2) == 0 {
		return nums1
	}

	var min, max []int

	if nums1[0] <= nums2[0] {
		min = nums1
		max = nums2
	} else {
		min = nums2
		max = nums1
	}

	minLen, maxLen := len(min), len(max)

	if max[0] >= min[minLen-1] {
		return append(min, max...)
	}

	i, j, k := 0, 0, 0
	result := make([]int, minLen+maxLen)
	for i<minLen && j <maxLen {
		if min[i] < max[j] {
			result[k] = min[i]
			i++
		} else {
			result[k] = max[j]
			j++
		}
		k++
	}

	if i != minLen{
		for i< minLen {
			result[k] = min[i]
			k++
			i++
		}
	}

	if j != maxLen{
		for j< maxLen {
			result[k] = max[j]
			k++
			j++
		}
	}
	return result
}

func getArrayMiddle(nums []int) float64 {
	fmt.Println(nums)
	l := len(nums)
	i := l / 2
	if l%2 == 0 {
		return float64(nums[i-1] + nums[i]) / 2
	} else {
		return float64(nums[i])
	}

}
