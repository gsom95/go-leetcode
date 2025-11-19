package keepmultiplyingfoundvaluesbytwo

import "slices"

func findFinalValue(nums []int, original int) int {
	slices.Sort(nums)
	for _, n := range nums {
		if n == original {
			original *= 2
		}
	}

	return original
}
