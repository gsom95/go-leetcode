package reversestringprefix

import "strings"

func reversePrefix(s string, k int) string {
	var sb strings.Builder
	sb.Grow(len(s))
	for i := k - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	sb.WriteString(s[k:])

	return sb.String()
}
