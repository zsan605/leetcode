package problems

import "sort"

func fourSum(nums []int, target int) [][]int {

	sort.Ints(nums)
	n := len(nums)
	var ret [][]int

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			four := n - 1
			temp := target - nums[first] - nums[second]
			for third := second + 1; third < n; third++ {
				if third > second+1 && nums[third] == nums[third-1] {
					continue
				}

				for third < four && nums[third]+nums[four] > temp {
					four--
				}

				if third == four {
					break
				}

				if nums[third]+nums[four] == temp {
					ret = append(ret, []int{nums[first], nums[second], nums[third], nums[four]})
				}
			}
		}
	}
	return ret
}
