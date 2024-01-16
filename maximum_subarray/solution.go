package maximumsubarray

import "math"

func maxSubArray(nums []int) int {
	sum := math.MinInt
	curSum := 0
	for _, n := range nums {
		curSum += n
		if n > curSum {
			curSum = n
		}
		if curSum > sum {
			sum = curSum
		}
	}

	return sum
}
