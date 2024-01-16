package bulbswitcher

import "math"

func bulbSwitch(n int) int {
	return int(math.Floor(math.Sqrt(float64(n))))
}
