package matchstickstosquare

import (
	"fmt"
	"sort"
)

func makesquare(matchsticks []int) bool {
	if len(matchsticks) < 4 {
		return false
	}
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})
	fmt.Println(matchsticks)
	sum := 0
	for _, v := range matchsticks {
		sum += v
	}
	if sum%4 != 0 {
		return false
	}

	return dfs(matchsticks, [4]int{})
}

func dfs(sticks []int, sums [4]int) bool {
	if len(sticks) == 0 {
		return sums[0] == sums[1] && sums[0] == sums[2] && sums[0] == sums[3]
	}

	elem := sticks[0]
	for i := 0; i < 4; i++ {
		sums[i] += elem
		if dfs(sticks[1:], sums) {
			return true
		}
		sums[i] -= elem
	}

	return false
}
