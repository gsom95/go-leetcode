package findthekthcharacterinstringgameii

import "math/bits"

func kthCharacter(k int64, operations []int) byte {
	ans := 0
	k--
	for i := bits.Len64(uint64(k)) - 1; i >= 0; i-- {
		if (k>>i)&1 == 1 {
			ans += operations[i]
		}
	}
	return byte('a' + (ans % 26))
}
