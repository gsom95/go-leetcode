package determine_if_string_halves_are_alike

import (
	"strings"
)

var vowels = map[rune]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
}

func halvesAreAlike(s string) bool {
	s = strings.ToLower(s)
	mid := len(s) / 2
	return countVowels(s[:mid]) == countVowels(s[mid:])
}

func countVowels(s string) int {
	count := 0
	for _, r := range s {
		if _, ok := vowels[r]; ok {
			count++
		}
	}
	return count
}
