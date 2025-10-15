package summaryranges

import "fmt"

func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	if len(nums) == 0 {
		return nil
	}

	left := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[left]+i-left {
			continue
		}
		right := i - 1
		if left == right {
			res = append(res, fmt.Sprintf("%d", nums[left]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[left], nums[right]))
		}
		left = i
	}
	if left < len(nums)-1 {
		res = append(res, fmt.Sprintf("%d->%d", nums[left], nums[len(nums)-1]))
	} else {
		res = append(res, fmt.Sprintf("%d", nums[left]))
	}

	return res
}
