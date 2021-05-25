package leetcode

//题目描述
//在一个长度为 n 的数组里的所有数字都在 0 到 n-1 的范围内。数组中某些数字是重复的，
//但不知道有几个数字是重复的，也不知道每个数字重复几次。请找出数组中任意一个重复的数字。
//
//Input:
//{2, 3, 1, 0, 2, 5}
//
//Output:
//2

func repeatNumberInArray(arr []int) int {

	if len(arr) <= 1 {
		return -1
	}

	for i, _ := range arr {
		for arr[i] != i {
			if arr[i] == arr[arr[i]] {
				return arr[i]
			}

			arr[i], arr[arr[i]] = arr[arr[i]], arr[i]
		}
	}
	return -1
}