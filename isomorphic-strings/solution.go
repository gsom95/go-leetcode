package isomorphic_strings

func isIsomorphic(s string, t string) bool {
	mapping := make(map[byte]byte)

	type seen struct{}
	usedChars := make(map[byte]seen)
	for i := range s {
		sc := s[i]
		ch := t[i]

		mappedChar, found := mapping[sc]
		if found {
			if ch == mappedChar {
				continue
			} else {
				return false
			}
		}

		if _, used := usedChars[ch]; used {
			return false
		}
		mapping[sc] = ch
		usedChars[ch] = seen{}
	}

	return true
}
