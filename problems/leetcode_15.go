package problems

import "sort"

func threeSum(nums []int) [][]int {

	sort.Ints(nums)

	n := len(nums)
	var ret [][]int

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		if nums[first] >= 0 {
			break
		}

		third := n - 1
		target := -1 * nums[first]
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			for second < third && nums[second]+nums[third] > target {
				third--
			}

			if second == third {
				break
			}

			if nums[second]+nums[third] == target {
				ret = append(ret, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return ret
}
