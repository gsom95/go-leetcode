package removeelement

func removeElement(nums []int, val int) int {
	insertTo := 0
	for _, num := range nums {
		if num != val {
			nums[insertTo] = num
			insertTo++
		}
	}
	return insertTo
}
