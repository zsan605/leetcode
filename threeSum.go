package leetcode

import (
	"sort"
	"strconv"
)


//[[-1,0,1],[-2,1,1],[-1,-1,2],[-2,0,2],[-3,1,2],[-2,-1,3],[-3,0,3],[-5,2,3],[-2,-2,4],[-5,1,4]]
//[[-1,0,1],[-2,1,1],[-1,-1,2],[-2,0,2],[-3,1,2],[-2,-1,3],[-3,0,3],[-5,2,3],[-2,-2,4],[-5,1,4],[-3,-1,4],,,,,,]

func threeSum(nums []int) [][]int {
	res := calcSubSet(nums)

	return res
}

func calcSubSet(nums []int) [][]int {
	sort.Ints(nums)
	var subSets [][]int
	var res [][]int
	tempMap := make(map[string]bool)
	for _, num := range nums {
		for _, subSet := range subSets {
			subSet = append(subSet, num)
			if len(subSet) == 3 {

				if checkSum(subSet) {
					key :=intArrToStr(subSet)
					_, ok := tempMap[key]
					if !ok {
						res = append(res, subSet)
						tempMap[key] =true
					}
				}
			} else {
				subSets = append(subSets, subSet)
			}
		}
		subSets = append(subSets, []int{num})
	}

	return res
}

func checkSum(nums []int) bool {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum == 0 {
		return true
	}
	return false
}

func intArrToStr(nums []int) string{

	str := ""
	for _, v := range nums {
		str += strconv.Itoa(v)
	}
	return str
}