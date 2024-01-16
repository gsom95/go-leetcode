package besttimetobuyandsellstock

import (
	"fmt"
	"math"
)

func MaxProfitI(prices []int) int {
	maxProfit := 0
	minPrice := math.MaxInt
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if (price - minPrice) > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}

func MaxProfitII(prices []int) int {
	pricesLen := len(prices)
	maxProfit := 0
	lowest := 0
	highest := 0
	for i := 0; i < pricesLen-1; {
		for ; i < pricesLen-1 && prices[i] >= prices[i+1]; i++ {
		}
		lowest = prices[i]

		for ; i < pricesLen-1 && prices[i] <= prices[i+1]; i++ {
		}
		highest = prices[i]
		maxProfit += highest - lowest
	}

	return maxProfit
}

func MaxProfitIII(prices []int) int {
	minPrice1 := math.MaxInt
	profit1 := 0
	minPrice2 := math.MaxInt
	profit2 := 0

	for _, curPrice := range prices {
		if curPrice < minPrice1 {
			minPrice1 = curPrice
		}
		if curPrice-minPrice1 > profit1 {
			profit1 = curPrice - minPrice1
		}
		if curPrice-profit1 < minPrice2 {
			minPrice2 = curPrice - profit1
		}
		if curPrice-minPrice2 > profit2 {
			profit2 = curPrice - minPrice2
		}
		fmt.Printf("minPrice1: %2v profit1: %2v minPrice2: %2v profit2: %2v\n", minPrice1, profit1, minPrice2, profit2)
	}

	return profit2
}

// with the help of state machine
// amount of states is amount_of_transactions*2
// - from s[0] we either go to s[1] (buying, i.e. spending currentPrice amount), or we stay in s[0] (no spending)
// - from s[1] we either go to s[2] (selling, i.e. receiving currentPrice amount )
// basically, s[i%2 == 0] is buying, i.e. we should subtract current price (we spend money)
// s[i%2 == 1] is selling, i.e. we should add current price (we get money)
func MaxProfitIV(k int, prices []int) int {
	if k == 0 || len(prices) < 2 {
		return 0
	}

	statesLen := 2 * k
	states := make([]int, statesLen)
	states[0] = -prices[0]
	for i := range states[1:] {
		states[i] = math.MinInt
	}

	for _, curPrice := range prices {
		states[0] = max(states[0], -curPrice)
		states[1] = max(states[1], states[0]+curPrice)
		for i := 2; i < statesLen; i += 2 {
			states[i] = max(states[i], states[i-1]-curPrice)
			states[i+1] = max(states[i+1], states[i]+curPrice)
		}
	}

	return max(0, states[statesLen-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
