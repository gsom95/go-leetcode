package constructkpalindromestrings

func canConstruct(s string, k int) bool {
	counter := [26]int{}
	for _, chr := range s {
		counter[chr-'a']++
	}

	oddAmount := 0
	for _, c := range counter {
		if c%2 == 1 {
			oddAmount++
		}
	}

	return len(s) >= k && oddAmount <= k
}
