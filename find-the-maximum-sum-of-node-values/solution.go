package findthemaximumsumofnodevalues

import "math"

func maximumValueSum(nums []int, k int, _ [][]int) int64 {
	sum := int64(0)
	count, positiveMin, negativeMax := 0, math.MaxInt, math.MinInt
	for _, num := range nums {
		sum += int64(num)
		netChange := num ^ k - num
		if netChange > 0 {
			positiveMin = min(positiveMin, netChange)
			sum += int64(netChange)
			count++
		} else {
			negativeMax = max(negativeMax, netChange)
		}
	}

	if count%2 == 0 {
		return sum
	}

	return max(sum-int64(positiveMin), sum+int64(negativeMax))
}
