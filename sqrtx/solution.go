package sqrtx

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	start := 0
	end := x
	mid := (start + end) / 2
	for start <= end {
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			end = mid - 1
		} else {
			start = mid + 1
		}
		mid = (start + end) / 2
	}

	return mid
}
