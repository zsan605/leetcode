package problems

/*
地址：https://leetcode-cn.com/problems/jump-game-ii/
给你一个非负整数数组nums ，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
假设你总是可以到达数组的最后一个位置。
[2,3,1,1,4]

贪心策略1：
1. 从后向前查找，找到最远一个可以跳到结尾的位置
2. 替换结尾位置为找到的最远位置，重复1，直至找到开头
*/
func jump(nums []int) int {

	// for i := len(nums) - 1; i >= 0; i-- {

	// 	ln := i - j
	// 	for j := 0; j < i; j++ {
	// 		if j+nums[j] >= i
	// 	}
	// }

	return 0
}
