package typeoftriangle

import "slices"

func triangleType(nums []int) string {
	slices.Sort(nums)
	switch {
	case nums[0]+nums[1] <= nums[2]:
		return "none"
	case nums[0] == nums[1] && nums[1] == nums[2]:
		return "equilateral"
	case nums[0] == nums[1] || nums[1] == nums[2]:
		return "isosceles"
	default:
		return "scalene"
	}
}
