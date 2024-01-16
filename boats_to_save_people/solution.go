package boatstosavepeople

import "sort"

func numRescueBoats(people []int, limit int) int {
	result := 0
	sort.Ints(people)

	l := 0
	r := len(people) - 1
	for l <= r {
		result++
		if people[l]+people[r] <= limit {
			l++
		}
		r--
	}

	return result
}
