package problems

import "math"

/*
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。

示例 1：

输入：x = 123
输出：321
示例 2：

输入：x = -123
输出：-321
示例 3：

输入：x = 120
输出：21
示例 4：

输入：x = 0
输出：0


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func reverse(x int) int {
	//if x >math.MaxInt32{
	//	return 0
	//}
	var ret = 0
	var flag = 1
	if x < 0 {
		flag = -1
	}
	for x%10 != 0 {
		ret = ret*10 + x%10
		x = x / 10

		if flag ==-1{
			if ret < (math.MaxInt32 +1) * flag{
				return  0
			}
		} else {
			if ret > math.MaxInt32{
				return 0
			}
		}
	}
	return ret
}