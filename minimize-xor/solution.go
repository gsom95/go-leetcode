package minimizexor

import "math/bits"

func minimizeXor(num1 int, num2 int) int {
	num1Bits := bits.OnesCount64(uint64(num1))
	num2Bits := bits.OnesCount64(uint64(num2))
	if num1Bits == num2Bits {
		return num1
	}

	x := 0
	for i := 63 - bits.LeadingZeros64(uint64(num1)); i >= 0 && num2Bits > 0; i-- {
		bit := num1 & (1 << i)
		if bit == 0 {
			continue
		}

		x ^= 1 << i
		num2Bits--
	}

	for i := 0; num2Bits > 0; i++ {
		bit := x & (1 << i)
		if bit != 0 {
			continue
		}

		x ^= 1 << i
		num2Bits--
	}

	return x
}
