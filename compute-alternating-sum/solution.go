package computealternatingsum

func alternatingSum(nums []int) int {
	sum := 0

	for i, n := range nums {
		if i%2 == 0 {
			sum += n
		} else {
			sum -= n
		}
	}

	return sum
}
