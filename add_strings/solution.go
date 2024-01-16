package add_strings

func addStrings(num1 string, num2 string) string {
	aLen := len(num1)
	bLen := len(num2)
	if aLen < bLen {
		num1, num2 = num2, num1
		aLen, bLen = bLen, aLen
	}

	result := make([]byte, 0, aLen+1)
	var carry byte
	for i := 0; i < aLen; i++ {
		var sum byte
		if i < bLen {
			sum = num1[aLen-1-i] - '0' + num2[bLen-1-i] - '0' + carry
		} else {
			sum = num1[aLen-1-i] - '0' + carry
		}

		if sum >= 10 {
			sum -= 10
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
