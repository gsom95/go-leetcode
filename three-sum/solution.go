package threesum

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	results := make([][]int, 0, len(nums))

	for i := range len(nums) - 2 {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target, left, right := -nums[i], i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				results = append(results, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
				continue
			}
			if sum > target {
				right--
				continue
			}
			left++
		}
	}

	return results
}
