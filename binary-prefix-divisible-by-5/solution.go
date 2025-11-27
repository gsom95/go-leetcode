package binaryprefixdivisibleby5

func prefixesDivBy5(nums []int) []bool {
	res := make([]bool, len(nums))
	curNum := 0
	for i, n := range nums {
		curNum = (curNum*2 + n) % 5 // modulo 5 to keep the number manageable
		if curNum == 0 {
			res[i] = true
		}
	}

	return res
}
