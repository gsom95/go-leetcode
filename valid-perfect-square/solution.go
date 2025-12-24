package validperfectsquare

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	left, right := 1, num
	for left <= right {
		mid := (left + right) / 2
		sqr := mid * mid
		if sqr == num {
			return true
		}

		if sqr < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
