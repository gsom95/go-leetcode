package countgoodtriplets

func countGoodTriplets(arr []int, a int, b int, c int) int {
	res := 0
	lenArr := len(arr)

	for i := 0; i < lenArr-2; i++ {
		for j := i + 1; j < lenArr-1; j++ {
			ij := arr[i] - arr[j]
			if ij < 0 {
				ij *= -1
			}

			if ij > a {
				continue
			}

			for k := j + 1; k < lenArr; k++ {
				jk := arr[j] - arr[k]
				if jk < 0 {
					jk *= -1
				}

				if jk > b {
					continue
				}

				ik := arr[i] - arr[k]
				if ik < 0 {
					ik *= -1
				}

				if ik > c {
					continue
				}

				res += 1
			}
		}
	}

	return res
}
