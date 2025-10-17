package adddigits

func addDigits(num int) int {
	if num == 0 {
		return 0
	}

	for num > 9 {
		sum := 0
		tempN := num
		for tempN > 0 {
			sum += tempN % 10
			tempN /= 10
		}
		num = sum
		sum = 0
	}

	return num
}
