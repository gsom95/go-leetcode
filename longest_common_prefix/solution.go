package longestcommonprefix

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var prefixBuilder strings.Builder
	maxLen := 0
	maxLenStr := ""
	for _, str := range strs {
		if len(str) > maxLen {
			maxLen = len(str)
			maxLenStr = str
		}
	}
	prefixBuilder.Grow(maxLen)

mainLoop:
	for i := 0; i < maxLen; i++ {
		curChar := maxLenStr[i]
		for _, str := range strs {
			if len(str) == 0 || (i+1) > len(str) || str[i] != curChar {
				break mainLoop
			}
		}
		prefixBuilder.WriteByte(curChar)
	}

	return prefixBuilder.String()
}
