package numberof1bits

func hammingWeight(n int) int {
	count := 0
	for range 31 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}
