package wordsubsets

func wordSubsets(words1 []string, words2 []string) []string {
	result := make([]string, 0, len(words1))

	words2Table := make(map[rune]int)
	for _, word2 := range words2 {
		counter := make(map[rune]int, len(word2))
		for _, wchr := range word2 {
			counter[wchr] += 1
		}

		for wchr, wcount := range counter {
			curCount, found := words2Table[wchr]
			if !found {
				words2Table[wchr] = wcount
				continue
			}

			if wcount > curCount {
				words2Table[wchr] = wcount
			}
		}
	}

	for _, word := range words1 {
		word1Counter := make(map[rune]int, len(word))
		for _, wchr := range word {
			word1Counter[wchr] += 1
		}

		isUnique := true
		for word2Chr, word2Count := range words2Table {
			word1Count, found := word1Counter[word2Chr]
			if !found || word1Count < word2Count {
				isUnique = false
				break
			}
		}
		if isUnique {
			result = append(result, word)
		}
	}

	return result
}
