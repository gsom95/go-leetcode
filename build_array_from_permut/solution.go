package buildarrayfrompermut

func buildArray(nums []int) []int {
	res := make([]int, len(nums))

	for i := range nums {
		res[i] = nums[nums[i]]
	}

	return res
}
