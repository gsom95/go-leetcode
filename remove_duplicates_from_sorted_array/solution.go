package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	insertTo := 0
	for curIdx := 0; curIdx < len(nums); curIdx++ {
		if nums[insertTo] != nums[curIdx] {
			insertTo++
			nums[insertTo] = nums[curIdx]
		}
	}
	return insertTo + 1
}
