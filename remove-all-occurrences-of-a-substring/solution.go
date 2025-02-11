package remove_all_occurrences_of_a_substring

import (
	"strings"
)

func removeOccurrences(s string, part string) string {
	index := strings.Index(s, part)
	for index != -1 {
		s = s[:index] + s[index+len(part):]
		index = strings.Index(s, part)
	}
	return s
}
