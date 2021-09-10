package problems

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return [][]int{nums}
	} else if len(nums) == 2 {
		return [][]int{nums, {nums[1], nums[0]}}
	}

	var ret [][]int
	for i := 0; i < len(nums); i++ {
		var arr []int
		arr = append(arr, nums[:i]...)
		arr = append(arr, nums[i+1:]...)
		temp := permute(arr)
		for j, _ := range temp {
			temp[j] = append([]int{nums[i]}, temp[j]...)
		}
		ret = append(ret, temp...)
	}

	return ret
}


