package minimumbitflipstoconvertnumber

func minBitFlips(start int, goal int) int {
	goal ^= start
	res := 0
	for i := range 64 {
		b := (goal >> i) & 1
		if b == 1 {
			res++
		}
	}

	return res
}
