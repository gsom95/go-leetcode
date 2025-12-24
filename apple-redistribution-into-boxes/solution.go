package appleredistributionintoboxes

import "slices"

func minimumBoxes(apple []int, capacity []int) int {
	res := 0

	slices.Sort(capacity)
	slices.Reverse(capacity)

	totalApples := 0
	for _, a := range apple {
		totalApples += a
	}

	for _, c := range capacity {
		totalApples -= c
		res++

		if totalApples <= 0 {
			break
		}
	}

	return res
}
