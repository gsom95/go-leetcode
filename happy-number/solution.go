package happy_number

func isHappy(n int) bool {
	seen := make(map[int]bool)
	for n != 1 {
		if seen[n] {
			return false
		}
		seen[n] = true
		n = sumOfSquares(n)
	}
	return true
}

func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
