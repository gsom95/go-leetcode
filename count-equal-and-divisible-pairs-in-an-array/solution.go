package countequalanddivisiblepairsinanarray

func countPairs(nums []int, k int) int {
	res := 0
	indices := make(map[int][]int)

	for i, num := range nums {
		curIndices := indices[num]
		curIndices = append(curIndices, i)
		indices[num] = curIndices
	}

	for _, curIndices := range indices {
		if len(curIndices) < 2 {
			continue
		}

		for i := 0; i < len(curIndices)-1; i++ {
			for j := i + 1; j < len(curIndices); j++ {
				if (curIndices[i]*curIndices[j])%k == 0 {
					res++
				}
			}
		}
	}

	return res
}
