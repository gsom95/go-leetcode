package successfulpairsofspellsandpotions

import (
	"math"
	"slices"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	pairs := make([]int, len(spells))
	lenp := len(potions)

	slices.Sort(potions)
	for i, spell := range spells {
		searching := int(math.Ceil(float64(success) / float64(spell)))
		pi := sort.SearchInts(potions, searching)
		pairs[i] = lenp - pi
	}

	return pairs
}
