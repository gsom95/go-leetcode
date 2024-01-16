package plus_one

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}

		digits[i] = 0
	}

	return append([]int{1}, digits...)
}
