package evaluatereversepolishnotation

import "strconv"

func evalRPN(tokens []string) int {
	var (
		stack = make([]int, 0, len(tokens))
	)

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			left, right := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, eval(token, left, right))
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[len(stack)-1]
}

func eval(operator string, left, right int) int {
	switch operator {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	}

	panic("at the disco")
}
