package decodethemessage

import "strings"

func decodeMessage(key string, message string) string {
	decodeMap := make(map[rune]rune, 26)

	curLetter := 'a'
	for _, keyChr := range key {
		if keyChr == ' ' {
			continue
		}
		if _, exists := decodeMap[keyChr]; exists {
			continue
		}
		decodeMap[keyChr] = curLetter
		curLetter++
		if curLetter > 'z' {
			break
		}
	}
	var decodedBuilder strings.Builder
	for _, msgChr := range message {
		if msgChr == ' ' {
			decodedBuilder.WriteRune(msgChr)
		} else {
			decodedBuilder.WriteRune(decodeMap[msgChr])
		}
	}
	return decodedBuilder.String()
}
