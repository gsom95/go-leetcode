package subsets

func subsets(nums []int) [][]int {
	res := make([][]int, 0, 1<<len(nums))
	cur := make([]int, 0, len(nums))

	var dfs func(start int)
	dfs = func(start int) {
		subset := make([]int, len(cur))
		copy(subset, cur)
		res = append(res, subset)

		for i := start; i < len(nums); i++ {
			cur = append(cur, nums[i])

			dfs(i + 1)

			cur = cur[:len(cur)-1]
		}
	}

	dfs(0)
	return res
}