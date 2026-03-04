package containerwithmostwater

func maxArea(height []int) int {
	maxArea := 0
	left, right := 0, len(height)-1
	for left < right {
		leftH, rightH := height[left], height[right]
		curArea := (right - left) * min(leftH, rightH)
		maxArea = max(maxArea, curArea)
		if leftH < rightH {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
