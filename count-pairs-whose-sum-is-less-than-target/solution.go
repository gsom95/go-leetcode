package countpairswhosesumislessthantarget

import "slices"

func countPairs(nums []int, target int) int {
	slices.Sort(nums)
	i, j := 0, len(nums)-1
	count := 0
	for i < j {
		if nums[i]+nums[j] < target {
			count += j - i
			i++
		} else {
			j--
		}
	}
	return count
}
