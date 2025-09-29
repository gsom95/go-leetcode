package reversebits

func reverseBits(n int) int {
	res := 0
	for i := 1; i < 32; i++ {
		temp := (n >> i) & 1
		res |= (temp << (31 - i))
	}
	return res
}
