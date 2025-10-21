package missingnumber

import "slices"

func missingNumber(nums []int) int {
	slices.Sort(nums)
	for i, n := range nums {
		if i != n {
			return i
		}
	}

	return len(nums)
}
