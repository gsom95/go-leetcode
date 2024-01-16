package containsduplicate

import "sort"

func containsDuplicate(nums []int) bool {
	sort.IntSlice(nums).Sort()
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}
	return false
}
