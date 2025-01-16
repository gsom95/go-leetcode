package bitwisexorofallpairings

func xorAllNums(nums1 []int, nums2 []int) int {
	var res int

	if len(nums2)%2 == 1 {
		for _, n := range nums1 {
			res ^= n
		}
	}
	if len(nums1)%2 == 1 {
		for _, n := range nums2 {
			res ^= n
		}
	}

	return res
}
