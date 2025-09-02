package findthenumberofwaystoplacepeoplei

import (
	"cmp"
	"slices"
)

func numberOfPairs(points [][]int) int {
	slices.SortFunc(points, func(a, b []int) int {
		if res1 := cmp.Compare(a[0], b[0]); res1 != 0 {
			return res1
		}
		return cmp.Compare(b[1], a[1])
	})

	var count int
	for i, firstPoint := range points {
	next:
		for j, secondPoint := range points[i+1:] {
			if firstPoint[0] > secondPoint[0] || firstPoint[1] < secondPoint[1] {
				continue
			}
			for _, tmpPoint := range points[i+1 : i+1+j] {
				if firstPoint[0] <= tmpPoint[0] && tmpPoint[0] <= secondPoint[0] &&
					firstPoint[1] >= tmpPoint[1] && tmpPoint[1] >= secondPoint[1] {
					continue next
				}
			}
			count++
		}
	}

	return count
}
