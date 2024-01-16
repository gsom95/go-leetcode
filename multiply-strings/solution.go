package multiply_strings

import (
	"slices"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	num1Len := len(num1)
	num2Len := len(num2)
	if num1Len < num2Len {
		num1, num2 = num2, num1
		num1Len, num2Len = num2Len, num1Len
	}

	num1Bytes := []byte(num1)
	num2Bytes := []byte(num2)
	slices.Reverse(num1Bytes)
	slices.Reverse(num2Bytes)

	result := make([]byte, 0, num1Len+num2Len)

	sum := 0
	for i := 0; i < num1Len; i++ {
		for j := 0; j < num2Len; j++ {
			sum += (int(num1Bytes[i]-'0') + 10*(num1Len-1-i)) * (int(num2Bytes[j]-'0') + 10*(num2Len-1-j))
		}
	}
	for sum > 0 {
		result = append(result, byte(sum%10)+'0')
		sum /= 10
	}
	slices.Reverse(result)

	return string(result)
}
