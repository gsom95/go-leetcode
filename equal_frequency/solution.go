package equalfrequency

func EqualFrequency(word string) bool {
	counter := make(map[rune]int)
	for _, letter := range word {
		counter[letter]++
	}
	if len(counter) == 1 {
		return true
	}
	

	return false
}
