package intersectionoftwoarrays

func intersection(nums1 []int, nums2 []int) []int {
	seen := map[int]struct{}{}
	for _, n := range nums1 {
		seen[n] = struct{}{}
	}
	res := make([]int, 0, len(seen))
	for _, n := range nums2 {
		if _, found := seen[n]; found {
			res = append(res, n)
			delete(seen, n)
		}
	}

	return res
}
