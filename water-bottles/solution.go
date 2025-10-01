package waterbottles

func numWaterBottles(numBottles int, numExchange int) int {
	drink := 0
	empty := 0
	for curN := numBottles; curN > 0; {
		drink += curN
		empty += curN

		curN = empty / numExchange
		empty = empty % numExchange
	}

	return drink
}
