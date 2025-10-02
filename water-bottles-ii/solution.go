package waterbottles

func numWaterBottles(numBottles int, numExchange int) int {
	res := numBottles
	for empty := numBottles; empty >= numExchange; numExchange++ {
		res++
		empty -= numExchange - 1
	}
	return res
}
