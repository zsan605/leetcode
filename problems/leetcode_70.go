package problems

// 递归
func climbStairs(n int) int {

	if n <= 2 {
		return n
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

// 迭代
func climbStairs_two(n int) int {

	var table = make([]int, n+1)
	table[0], table[1] = 1, 2
	for i := 2; i <= n; i++ {
		table[i] = table[i-1] + table[i-2]
	}
	return table[n-1]
}
