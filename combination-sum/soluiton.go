package combinationsum

import "slices"

func combinationSum(candidates []int, target int) [][]int {
	slices.Sort(candidates)

	var ans [][]int
	var curPath []int

	var dfs func(start int, sum int)
	dfs = func(start, sum int) {
		if sum == target {
			set := make([]int, len(curPath))
			copy(set, curPath)
			ans = append(ans, set)
			return
		}

		for i := start; i < len(candidates); i++ {
			num := candidates[i]

			if sum+num > target {
				break
			}

			curPath = append(curPath, num)
			dfs(i, sum+num)

			curPath = curPath[:len(curPath)-1]
		}
	}

	dfs(0, 0)
	return ans
}
