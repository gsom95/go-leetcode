package firstuniquecharacterinstring

func firstUniqChar(s string) int {
	counter := make([]int, 26)
	for _, chr := range s {
		counter[chr-'a']++
	}
	for i, chr := range s {
		if amnt := counter[chr-'a']; amnt == 1 {
			return i
		}
	}
	return -1
}
