package buildanarraywithstackoperations

func buildArray(target []int, n int) []string {
	res := make([]string, 0, len(target))

	tIdx := 0
	for i := 1; i <= n && tIdx < len(target); i++ {
		res = append(res, "Push")
		if i != target[tIdx] {
			res = append(res, "Pop")
			continue
		}
		tIdx++
	}

	return res
}
