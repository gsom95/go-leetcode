package intersectionoftwoarrays

func intersection(nums1 []int, nums2 []int) []int {
	seen := map[int]int{}
	for _, n := range nums1 {
		seen[n]++
	}
	res := make([]int, 0, len(seen))
	for _, n := range nums2 {
		if amount := seen[n]; amount > 0 {
			res = append(res, n)
			seen[n]--
		}
	}

	return res
}
