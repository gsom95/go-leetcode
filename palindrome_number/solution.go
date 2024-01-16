package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverse := 0
	temp := x
	for temp > 0 {
		reverse = (reverse * 10) + temp%10
		temp /= 10
	}

	return x == reverse
}
