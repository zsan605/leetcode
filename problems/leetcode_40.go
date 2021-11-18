package problems

import "sort"

func combinationSum2(candidates []int, target int) [][]int {

	if len(candidates) == 0 || target == 0 {
		return [][]int{}
	}

	sort.Ints(candidates)

	var l = len(candidates)
	return combinationSum2_dfs(candidates, 0, l, target)
}

func combinationSum2_dfs(candidates []int, pos, l, target int) [][]int {

	var ret [][]int
	if target == 0 || pos+1 == l {
		return [][]int{}
	}

	for i := pos; i < l; i++ {
		if candidates[pos] == target {
			ret = append(ret, []int{candidates[pos]})
		} else if candidates[pos] < target {
			tempRet := combinationSum2_dfs(candidates, pos+1, l, target-candidates[pos])
			for _, val := range tempRet {
				ret = append(ret, append(val, candidates[pos]))
			}
		}
	}
	return ret
}
