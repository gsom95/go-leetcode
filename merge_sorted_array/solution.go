package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, m+n)
	resInsert := 0
	first := 0
	second := 0
	for first < m && second < n {
		if nums1[first] < nums2[second] {
			res[resInsert] = nums1[first]
			first++
		} else {
			res[resInsert] = nums2[second]
			second++
		}
		resInsert++
	}
	if first < m {
		copy(res[resInsert:], nums1[first:])
	}
	if second < n {
		copy(res[resInsert:], nums2[second:])
	}

	copy(nums1, res)
}
