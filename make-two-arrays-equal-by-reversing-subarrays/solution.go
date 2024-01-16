package make_two_arrays_equal_by_reversing_subarrays

import (
	"sort"
)

func canBeEqual(target []int, arr []int) bool {
	targetValues := make(map[int]int)
	for _, v := range target {
		targetValues[v]++
	}

	for _, v := range arr {
		if tvCount := targetValues[v]; tvCount > 0 {
			targetValues[v]--
		} else {
			return false
		}
	}

	for _, v := range targetValues {
		if v != 0 {
			return false
		}
	}

	return true
}

func canBeEqualSorted(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)

	for i := range target {
		if target[i] != arr[i] {
			return false
		}
	}

	return true
}
