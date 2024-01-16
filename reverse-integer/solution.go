package reverseinteger

import "math"

func reverse(x int) int {
	isNegative := false

	if x < 0 {
		x *= -1
		isNegative = true
	}

	var res int
	for x > 9 {
		num := x % 10
		res = res*10 + num
		x /= 10
	}
	res = res*10 + x
	if res >= math.MaxInt32 || res <= math.MinInt32 {
		return 0
	}

	if isNegative {
		res *= -1
	}

	return res
}
