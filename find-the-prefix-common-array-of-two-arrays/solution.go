package findtheprefixcommonarrayoftwoarrays

func findThePrefixCommonArray(A []int, B []int) []int {
	var (
		aFound, bFound uint64
		res            = make([]int, len(A))
	)

	for i := 0; i < len(A); i++ {
		aFound |= 1 << A[i]
		bFound |= 1 << B[i]
		commonFound := aFound & bFound

		commons := 0
		for j := 0; j < 64; j++ {
			if commonFound&(1<<j) > 0 {
				commons++
			}
		}
		res[i] = commons
	}

	return res
}
