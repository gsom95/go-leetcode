package countmonobitintegers

func countMonobit(n int) int {
	res := 1

	for i := 2; (i - 1) <= n; i *= 2 {
		res++
	}

	return res
}
