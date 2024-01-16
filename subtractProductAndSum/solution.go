package subtractProductAndSum

func subtractProductAndSum(n int) int {
	sum := 0
	prod := 1
	for n > 0 {
		num := n % 10
		sum += num
		prod *= num

		n = n / 10
	}

	return prod - sum
}
