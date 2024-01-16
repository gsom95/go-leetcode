package cutthesticks

import (
	"sort"
)

func cutTheSticks(arr []int) []int {
	result := make([]int, 0, len(arr))

	sort.Ints(arr)
	arrLen := len(arr)

	i := 0
	leastChosen := 0
	for i < arrLen {
		result = append(result, arrLen-i)
		leastChosen = arr[i]

		for j := i; j < arrLen; j++ {
			arr[j] -= leastChosen
			if arr[j] == 0 {
				i++
			}
		}

	}

	return result
}
