package reversewordsinastring3

import (
	"strings"
)

func reverseWords(s string) string {
	return reverseWordsWithNewBulder(s)
}

func reverseWordsWithReset(s string) string {
	words := strings.Split(s, " ")
	var wordBuilder strings.Builder
	for wi, word := range words {
		wordBuilder.Reset()
		wordBuilder.Grow(len(word))

		for i := len(word) - 1; i >= 0; i-- {
			wordBuilder.WriteByte(word[i])
		}
		words[wi] = wordBuilder.String()
	}

	return strings.Join(words, " ")
}

func reverseWordsWithNewBulder(s string) string {
	words := strings.Split(s, " ")
	for wi, word := range words {
		var wordBuilder strings.Builder
		wordBuilder.Grow(len(word))

		for i := len(word) - 1; i >= 0; i-- {
			wordBuilder.WriteByte(word[i])
		}
		words[wi] = wordBuilder.String()
	}

	return strings.Join(words, " ")
}
