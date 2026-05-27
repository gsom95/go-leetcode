package howmanynumbersaresmallerthanthecurrentnumber

func smallerNumbersThanCurrent(nums []int) []int {
	count := make([]int, 101) // 0 <= n <= 100
	for _, n := range nums {
		count[n]++
	}

	for i := range 100 {
		count[i+1] += count[i]
	}

	ans := make([]int, len(nums))
	for i, n := range nums {
		if n == 0 {
			continue
		}
		ans[i] = count[n-1]
	}

	return ans
}
