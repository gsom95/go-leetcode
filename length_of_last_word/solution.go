package lengthoflastword

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	words := strings.Split(s, " ")
	return len(words[len(words)-1])
}
