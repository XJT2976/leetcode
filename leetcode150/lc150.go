package leetcode150

import "strconv"

func EvalRPN(tokens []string) int {
	stack := make([]string, len(tokens))
	j := 0

	for _, s := range tokens {
		switch s {
		case "+", "-", "*", "/":
			num2, _ := strconv.Atoi(stack[j-1])
			num1, _ := strconv.Atoi(stack[j-2])
			if s == "+" {
				stack[j-2] = strconv.Itoa(num1 + num2)
			} else if s == "-" {
				stack[j-2] = strconv.Itoa(num1 - num2)
			} else if s == "*" {
				stack[j-2] = strconv.Itoa(num1 * num2)
			} else {
				stack[j-2] = strconv.Itoa(num1 / num2)
			}
			j--
		default:
			stack[j] = s
			j++
		}
	}

	ret, _ := strconv.Atoi(stack[0])
	return ret
}
