package problems

/*
地址：https://leetcode-cn.com/problems/first-bad-version/
思路：
1. 二分查找
2. 定义首尾指针
3. 找到中间位置，判断中间位置是否为badversion
4. 如果是，则尾指针为，中间位置（第一个坏版本在首指针到中间位置的区间里面）
5. 如果不是，则首指针为，中间位置+1
6. 知道尾指针<= 首指针
7. 此时返回首指针
*/
func firstBadVersion(n int) int {
	i, j := 1, n

	var m int

	for i < j {
		m = (i + j) / 2
		if isBadVersion(m) {
			j = m
		} else {
			i = m + 1
		}
	}
	return i
}

func isBadVersion(version int) bool {
	return true
}
