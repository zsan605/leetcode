package problems

func generateParenthesis(n int) []string {

	var ret []string
	generateParenthesisHelper("", n, n, &ret)
	return ret
}


func generateParenthesisHelper(str string, left, right int, ret *[]string) {

	if left == 0 && right == 0 {
		*ret = append(*ret, str)
		return
	}
	if left == right {
		generateParenthesisHelper(str+"(", left-1, right,  ret)
	} else if left < right {
		if left > 0 {
			generateParenthesisHelper(str+"(", left-1, right, ret)
		}
		generateParenthesisHelper(str+")", left, right-1, ret)

	}
}
