package removingstarsfromastring

func removeStars(s string) string {
	stack := make([]byte, 0, len(s))

	for i := range s {
		if s[i] != '*' {
			stack = append(stack, s[i])
			continue
		}
		stack = stack[:len(stack)-1]
	}

	return string(stack)
}
