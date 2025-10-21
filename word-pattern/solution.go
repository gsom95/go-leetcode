package wordpattern

import "strings"

func wordPattern(pattern string, s string) bool {
	splitWords := strings.Split(s, " ")
	if len(pattern) != len(splitWords) {
		return false
	}

	pToW := make(map[rune]string, len(splitWords))

	for i, pattern := range pattern {
		word := splitWords[i]
		expWord, exists := pToW[pattern]
		if exists {
			if expWord != word {
				return false
			}
		}
		for p, w := range pToW {
			if w == word && p != pattern {
				return false
			}
		}
		pToW[pattern] = word
	}

	return true
}
