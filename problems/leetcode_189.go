package problems

/*
地址：https://leetcode-cn.com/problems/rotate-array/
*/

func rotate(nums []int, k int) {
	k = k % len(nums)
	tempNums := nums[0:k]
	for i := 1; i <= k; i++ {

		var j int
		temp := tempNums[i-1]
		for j = i - 1; j+k <= len(nums)-1; j = j + k {
			temp, nums[j+k] = nums[j+k], temp
		}

		nums[(j+k)%len(nums)] = temp
	}
}

/*
思路：反转数组
1. 将数组反转
2. 将反转后的前k个反转
3. 将反转后的后n-k个反转
*/
func rotate1(nums []int, k int) {
	k = k % len(nums)
	rotateReverse(nums)
	rotateReverse(nums[0:k])
	rotateReverse(nums[k:])
}

func rotateReverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

//func rotate(nums []int, k int) {
//	k = k % len(nums)
//
//	for i := 1; i <= k; i++ {
//		temp := nums[i-1]
//		j := i - 1 + k
//		for ; j < len(nums)-1; j = j + k {
//			nums[j], temp = temp, nums[j]
//		}
//		nums[j%len(nums)] = temp
//	}
//}
