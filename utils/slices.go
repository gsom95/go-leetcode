package utils

// EqualIntSlices checks if two int slices are equal.
func EqualIntSlices(left, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for i, lv := range left {
		if lv != right[i] {
			return false
		}
	}

	return true
}

func EqualInt32Slices(left, right []int32) bool {
	if len(left) != len(right) {
		return false
	}
	for i, lv := range left {
		if lv != right[i] {
			return false
		}
	}

	return true
}

func EqualBoolSlices(left, right []bool) bool {
	if len(left) != len(right) {
		return false
	}

	for i, lv := range left {
		if lv != right[i] {
			return false
		}
	}

	return true
}
