package zeroarraytransformationi

func isZeroArray(nums []int, queries [][]int) bool {
	diff := make([]int, len(nums)+1)
	for _, q := range queries {
		l, r := q[0], q[1]
		diff[l]++
		diff[r+1]--
	}

	curr := 0
	for i := 0; i < len(nums); i++ {
		curr += diff[i]
		if curr < nums[i] {
			return false
		}
		nums[i] -= curr
		if nums[i] < 0 {
			nums[i] = 0
		}
	}

	for _, v := range nums {
		if v != 0 {
			return false
		}
	}
	return true
}
