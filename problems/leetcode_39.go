package problems

func combinationSum(candidates []int, target int) [][]int {
	var ret [][]int

	if len(candidates) == 0 {
		return ret
	}
	for i, val := range candidates {
		if target == val {
			ret = append(ret, []int{val})
		} else if target > val {
			temp := combinationSum(candidates[i:], target-val)
			for _, a := range temp {
				ret = append(ret, append(a, val))
			}
		}
	}

	return ret
}
