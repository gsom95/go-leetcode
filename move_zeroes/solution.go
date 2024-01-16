package movezeroes

func moveZeroes(nums []int) {
	insertTo := 0
	for i, n := range nums {
		if n != 0 {
			nums[insertTo], nums[i] = nums[i], nums[insertTo]
			insertTo++
		}
	}
}
