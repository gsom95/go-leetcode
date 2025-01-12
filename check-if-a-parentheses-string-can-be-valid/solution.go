package checkifaparenthesesstringcanbevalid

func canBeValid(s string, locked string) bool {
	if len(s)%2 == 1 {
		return false
	}

	openP := make([]int, 0, len(s))
	unlocked := make([]int, 0, len(s))

	for i, p := range s {
		if locked[i] == '0' {
			unlocked = append(unlocked, i)
			continue
		}
		if p == '(' {
			openP = append(openP, i)
			continue
		}

		lenOpenP := len(openP)
		if lenOpenP > 0 {
			openP = openP[:lenOpenP-1]
			continue
		}

		lenUnlocked := len(unlocked)
		if lenUnlocked > 0 {
			unlocked = unlocked[:lenUnlocked-1]
			continue
		}

		return false
	}

	for lenOpenP, lenUnlocked := len(openP), len(unlocked); lenOpenP > 0 && lenUnlocked > 0 && openP[lenOpenP-1] < unlocked[lenUnlocked-1]; lenOpenP, lenUnlocked = len(openP), len(unlocked) {
		openP = openP[:lenOpenP-1]
		unlocked = unlocked[:lenUnlocked-1]
	}

	return len(openP) == 0
}
