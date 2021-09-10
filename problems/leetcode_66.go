package problems

func plusOne(digits []int) []int {

	var flag = 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + flag
		flag = 0
		if digits[i] < 10 {
			break
		}

		digits[i] = digits[i] % 10
		flag = 1
	}

	if flag == 1 {
		return append([]int{1}, digits...)
	}

	return digits
}
