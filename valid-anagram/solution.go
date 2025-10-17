package validanagram

func isAnagram(s string, t string) bool {
	sc := make(map[rune]int, len(s))
	for _, sChr := range s {
		sc[sChr]++
	}

	for _, tChr := range t {
		if _, ok := sc[tChr]; !ok {
			return false
		}
		sc[tChr]--
	}

	for _, c := range sc {
		if c != 0 {
			return false
		}
	}

	return true
}
