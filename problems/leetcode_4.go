package problems

/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
示例 3：

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
示例 4：

输入：nums1 = [], nums2 = [1]
输出：1.00000
示例 5：

输入：nums1 = [2], nums2 = []
输出：2.00000


提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106


进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
解题思路：
	用下标给是两个数组排序
	1. 如果num1[max] < num2[0], 可拼接num1 num
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	i, j := 0, 0
	var nums3 []int
	for i < l1 && j < l2 {
		if nums1[i] < nums2[j] {
			nums3 = append(nums3, nums1[i])
			i++
		} else {
			nums3 = append(nums3, nums2[j])
			j++
		}
	}

	for i < l1 {
		nums3 = append(nums3, nums1[i])
		i++
	}

	for j < l2 {
		nums3 = append(nums3, nums2[j])
		j++
	}

	l3 := len(nums3)
	mid, left := l3/2, l3%2
	if left == 0 {
		return float64((nums3[mid-1] + nums3[mid])) / 2
	} else {
		return float64(nums3[mid])
	}
}

// func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
// 	l1, l2, l3 := len(nums1), len(nums2), 0
// 	i, j := 0, 0
// 	leftValue, rightValue := 0, 0

// 	mid, left := (l1+l2)/2, (l1+l2)%2

// 	for i < l1 && j < l2 {
// 		if nums1[i] < nums2[j] {
// 			if left != 0 && i+j == mid {
// 				return float64(nums1[i])
// 			}

// 			if left == 0 && i+j == mid -1 {

// 			}
// 			i++
// 		} else {
// 			if left != 0 && i+j == mid {
// 				return float64(nums1[i])
// 			}
// 			j++
// 		}

// 	}

// 	for i < l1 {
// 		i++
// 	}

// 	for j < l2 {
// 		j++
// 	}

// 	mid, left := l3/2, l3%2
// 	if left == 0 {
// 		return float64((nums3[mid-1] + nums3[mid])) / 2
// 	} else {
// 		return float64(nums3[mid])
// 	}
// }
