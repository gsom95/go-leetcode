package removelettertoequalizefrequency

import (
	"maps"
	"slices"
)

func equalFrequency(word string) bool {
	count := make(map[rune]int, len(word))
	for _, ch := range word {
		count[ch]++
	}

	freqs := make(map[int]int, len(count))
	for _, f := range count {
		freqs[f]++
	}

	if len(freqs) > 2 {
		return false
	}

	if len(freqs) == 1 {
		_, found := freqs[1]
		if found {
			return true
		}
		for _, v := range freqs {
			if v == 1 {
				return true
			}
		}
		return false
	}

	finalFreqs := slices.Collect(maps.Keys(freqs))
	slices.Sort(finalFreqs)
	f0, f1 := finalFreqs[0], finalFreqs[1]
	c0, c1 := freqs[f0], freqs[f1]
	if f1-f0 != 1 {
		return f0 == 1 && c0 == 1
	}

	return c1 == 1 || (f0 == 1 && c0 == 1)
}
