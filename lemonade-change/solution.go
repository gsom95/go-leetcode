package lemonadechange

func lemonadeChange(bills []int) bool {
	var five, ten int

	for _, bill := range bills {
		switch bill {
		case 5:
			five++
		case 10:
			ten++
		}

		change := bill - 5
		for change > 0 && (ten > 0 || five > 0) {
			if change-10 >= 0 && ten > 0 {
				ten--
				change -= 10
				continue
			}
			if change-5 >= 0 && five > 0 {
				five--
				change -= 5
				continue
			}

			return false
		}
		if change > 0 {
			return false
		}
	}

	return true
}
