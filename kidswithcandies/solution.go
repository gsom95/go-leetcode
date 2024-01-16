package kidswithcandies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))

	max := 0
	for _, candy := range candies {
		if candy > max {
			max = candy
		}
	}

	for i, kids := range candies {
		if (kids + extraCandies) >= max {
			result[i] = true
		}
	}

	return result
}
