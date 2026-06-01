package validpalindrome

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	runes := []rune(s)

	for left, right := 0, len(runes)-1; left < right; {
		if (runes[left] < 'a' || runes[left] > 'z') && (runes[left] < '0' || runes[left] > '9') {
			left++
			continue
		}
		if (runes[right] < 'a' || runes[right] > 'z') && (runes[right] < '0' || runes[right] > '9') {
			right--
			continue
		}
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}

	return true
}
