package reversedegreeofastring

func reverseDegree(s string) int {
	sum := 0
	for i, chr := range s {
		chr := int(chr)
		sum += (i + 1) * (26 - chr + 'a')
	}

	return sum
}
