package fruitintobaskets

func totalFruit(fruits []int) int {
	if len(fruits) == 1 {
		return 1
	}

	maxNum := 1
	tempMax := maxNum

	firstType := fruits[0]
	firstTypeStart := 0
	secondType := -1
	secondTypeStart := -1

	for i := 1; i < len(fruits); i++ {
		fruitType := fruits[i]

		if fruitType == firstType || fruitType == secondType {
			tempMax++
			if fruitType == firstType && fruits[i-1] == secondType {
				firstTypeStart = i
			} else if fruitType == secondType && fruits[i-1] == firstType {
				secondTypeStart = i
			}
			continue
		}

		if secondType == -1 {
			secondType = fruitType
			secondTypeStart = i
			tempMax++
			continue
		}

		if tempMax > maxNum {
			maxNum = tempMax
		}

		if secondTypeStart > firstTypeStart {
			firstType = secondType
			firstTypeStart = secondTypeStart
		}

		secondType = fruitType
		secondTypeStart = i
		tempMax = i - firstTypeStart + 1
	}

	if tempMax > maxNum {
		maxNum = tempMax
	}

	return maxNum
}
