package eliminate_maximum_number_of_monsters

import (
	"sort"
)

func eliminateMaximum(dist []int, speed []int) int {
	arrivals := make([]float64, len(dist))
	for i := range dist {
		arrivals[i] = float64(dist[i]) / float64(speed[i])
	}

	sort.Float64s(arrivals)
	n := 0
	for i := range arrivals {
		if arrivals[i] <= float64(i) {
			break
		}
		n++
	}

	return n
}
