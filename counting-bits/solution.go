package counting_bits

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		result[i] = countBit(i)
	}
	return result
}

func countBit(n int) int {
	count := 0
	for n > 0 {
		n &= n - 1
		count++
	}
	return count
}
