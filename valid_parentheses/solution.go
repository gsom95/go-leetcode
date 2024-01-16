package valid_parentheses

func isValid(s string) bool {
	var stack []rune

	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}

	for _, cur := range s {
		if cur == '(' || cur == '[' || cur == '{' {
			stack = append(stack, cur)
		} else {
			if len(stack) == 0 {
				return false
			}

			prev := stack[len(stack)-1]
			if prev == '(' && cur == ')' || prev == '[' && cur == ']' || prev == '{' && cur == '}' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
