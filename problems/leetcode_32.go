package problems

func longestValidParentheses(s string) int {

	var num, top, max = 0, -1, 0
	var stack = make([]byte, len(s))

	if s == "" {
		return 0
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			top++
			stack[top] = s[i]
			continue
		}

		if s[i] == ')' {

			if top == -1 {
				if max < num {
					max = num
					num = 0
				}
				continue
			}
			if stack[top] == '(' {
				num += 2
				top--
			}
		}
	}

	if max < num {
		max = num
		num = 0
	}

	return max

}

//func longestValidParentheses(s string) int {
//
//	var num, top, max = 0, -1, 0
//	var stack = make([]byte, len(s))
//
//	if s == "" {
//		return 0
//	}
//
//	for i := 0; i < len(s); i++ {
//		if top == -1 || s[i] == '(' {
//			top++
//			stack[top] = s[i]
//			continue
//		}
//
//		if s[i] == ')' {
//			if stack[top] == '(' {
//				num += 2
//				top--
//			} else {
//				top++
//				stack[top] = s[i]
//				if max < num {
//					max = num
//				}
//			}
//			continue
//		}
//	}
//
//	return num
//
//}
