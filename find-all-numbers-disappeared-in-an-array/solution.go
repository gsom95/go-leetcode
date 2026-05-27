package findallnumbersdisappearedinanarray

import "slices"

func findDisappearedNumbers(nums []int) []int {
	slices.Sort(nums)
	lenN := len(nums)
	cur := 1
	i := 0

	ans := make([]int, 0, lenN)
	for cur <= lenN && i < lenN {
		curN := nums[i]
		if curN == cur {
			cur++
			i++
			continue
		}
		if curN < cur {
			i++
			continue
		}

		ans = append(ans, cur)
		cur++
	}

	for ; cur <= lenN; cur++ {
		ans = append(ans, cur)
	}

	return ans
}
