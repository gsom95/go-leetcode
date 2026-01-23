package thirdmaximumnumber

import (
	"maps"
	"slices"
)

func thirdMax(nums []int) int {
	unique := map[int]struct{}{}
	for n := range slices.Values(nums) {
		unique[n] = struct{}{}
	}

	uslice := slices.Collect(maps.Keys(unique))
	slices.Sort(uslice)
	if len(uslice) == 1 {
		return uslice[0]
	}
	if len(uslice) == 2 {
		return uslice[1]
	}
	return uslice[len(uslice)-3]
}
