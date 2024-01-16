package leftandrightsumdifferences

// src: https://leetcode.com/problems/left-and-right-sum-differences/

func leftRigthDifference(nums []int) []int {
	if len(nums) == 1 {
		return []int{0}
	}

	result := make([]int, len(nums))
	for i := range nums {
		leftSum := sum(nums[:i])
		rightSum := sum(nums[i+1:])
		s := leftSum - rightSum
		if s < 0 {
			s *= -1
		}
		result[i] = s
	}
	return result
}

func sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for _, n := range nums {
		sum += n
	}

	return sum
}
