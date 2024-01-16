package maxconsecutiveones

func findMaxConsecutiveOnes(nums []int) int {
	maxOnes := 0
	left := 0
	zeroes := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroes++
		}

		for zeroes > 0 {
			if nums[left] == 0 {
				zeroes--
			}
			left++
		}
		maxOnes = max(maxOnes, right-left+1)
	}

	return maxOnes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
