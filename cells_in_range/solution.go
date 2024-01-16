package cellsinrange

import "fmt"

func cellsInRange(s string) []string {
	startLetter := s[0]
	startIndex := s[1]
	endLetter := s[3]
	endIndex := s[4]
	res := make([]string, 0, (endLetter-startLetter+1)*(endIndex-startIndex+1))
	for letter := startLetter; letter <= endLetter; letter++ {
		for index := startIndex; index <= endIndex; index++ {
			res = append(res, fmt.Sprintf("%c%c", letter, index))
		}
	}
	return res
}
