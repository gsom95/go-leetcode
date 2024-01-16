package maximumaveragesubarray1

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	div := float64(k)
	avg := float64(sum) / div

	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		tempAvg := float64(sum) / div
		if tempAvg > avg {
			avg = tempAvg
		}
	}

	return avg
}
