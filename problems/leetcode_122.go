package problems

func maxProfit2(prices []int) int {
	dp_0, dp_1 := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		dp_0, dp_1 = max(dp_0, dp_1+prices[i]), max(dp_1, dp_0-prices[i])

	}
	return dp_0
}
