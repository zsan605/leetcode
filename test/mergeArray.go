package test

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

 

说明：

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 

示例：

输入：
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出：[1,2,2,3,5,6]
 

提示：

-10^9 <= nums1[i], nums2[i] <= 10^9
nums1.length == m + n
nums2.length == n


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func merge(nums1 []int, m int, nums2 []int, n int) {

	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

	for n > 0 {
		nums1[n-1] = nums2[n-1]
		n--
	}

}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	for i<m && j < n {
		if nums1[i] >nums2[j] {
			nums1[i], nums2[j] = nums2[j], nums1[i]
			i++
		}else {

		}
	}
}