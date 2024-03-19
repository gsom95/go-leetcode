package find_words_containing_character

import (
	"strings"
)

func findWordsContaining(words []string, x byte) []int {
	res := make([]int, 0, len(words))
	for i, word := range words {
		if strings.ContainsRune(word, rune(x)) {
			res = append(res, i)
		}
	}

	return res
}
