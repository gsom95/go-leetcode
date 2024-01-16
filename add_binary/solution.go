package addbinary

import (
	"math/big"
)

func addBinary(a string, b string) string {
	aLen := len(a)
	var aNum big.Int
	for i := 0; i < aLen; i++ {
		if a[i] == '0' {
			aNum.SetBit(&aNum, aLen-1-i, 0)
		} else {
			aNum.SetBit(&aNum, aLen-1-i, 1)
		}
	}

	bLen := len(b)
	var bNum big.Int
	for i := 0; i < bLen; i++ {
		if b[i] == '0' {
			bNum.SetBit(&bNum, bLen-1-i, 0)
		} else {
			bNum.SetBit(&bNum, bLen-1-i, 1)
		}
	}

	return aNum.Add(&aNum, &bNum).Text(2)
}

func addBinaryPlain(a, b string) string {
	aLen := len(a)
	bLen := len(b)
	if aLen < bLen {
		a, b = b, a
		aLen, bLen = bLen, aLen
	}

	result := make([]byte, 0, aLen+1)
	var carry byte
	for i := 0; i < aLen; i++ {
		var sum byte
		if i < bLen {
			sum = a[aLen-1-i] - '0' + b[bLen-1-i] - '0' + carry
		} else {
			sum = a[aLen-1-i] - '0' + carry
		}

		if sum >= 2 {
			sum -= 2
			carry = 1
		} else {
			carry = 0
		}

		result = append(result, sum+'0')
	}

	if carry == 1 {
		result = append(result, '1')
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
