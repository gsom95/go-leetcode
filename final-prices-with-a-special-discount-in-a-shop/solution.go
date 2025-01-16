package finalpriceswithaspecialdiscountinashop

func finalPrices(prices []int) []int {
	result := make([]int, len(prices))

	pricesStack := make([]int, 0, len(prices))
	for i := len(prices) - 1; i >= 0; i-- {
		price := prices[i]

		for l := len(pricesStack) - 1; l >= 0; l-- {
			nextPrice := pricesStack[l]

			if nextPrice <= price {
				price -= nextPrice
				break
			}

			pricesStack = pricesStack[:l]
		}

		result[i] = price
		pricesStack = append(pricesStack, prices[i])
	}

	return result
}
