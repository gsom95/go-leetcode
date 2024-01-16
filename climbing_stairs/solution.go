package climbingstairs

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	a := 1
	b := 2
	for i := 3; i <= n; i++ {
		t := a+b
		a = b
		b = t
	}
	return b
}
