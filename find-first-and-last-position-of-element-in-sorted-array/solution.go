package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left, right := 0, len(nums)-1
	first, last := -1, -1

	// search first first
	for left <= right {
		mid := (left + right) / 2
		n := nums[mid]
		if n == target {
			first = mid
			right = mid - 1
			continue
		}
		if n < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if first == -1 {
		return []int{-1, -1}
	}

	// search last
	left, right = 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		n := nums[mid]
		if n == target {
			last = mid
			left = mid + 1
			continue
		}
		if n < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return []int{first, last}
}
