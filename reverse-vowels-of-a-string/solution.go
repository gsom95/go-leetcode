package reversevowelsofastring

import "strings"

func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	vowelsIndexes := make([]int, 0, len(s))

	for i, chr := range s {
		if strings.ContainsRune(vowels, chr) {
			vowelsIndexes = append(vowelsIndexes, i)
		}
	}

	sArr := []rune(s)
	lvi := len(vowelsIndexes)
	for i := 0; i < lvi/2; i++ {
		sArr[vowelsIndexes[i]], sArr[vowelsIndexes[lvi-1-i]] = sArr[vowelsIndexes[lvi-1-i]], sArr[vowelsIndexes[i]]
	}

	return string(sArr)
}
