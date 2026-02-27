package reversestringii

func reverseStr(s string, k int) string {
	sBytes := []byte(s)[:len(s):len(s)]
	if len(sBytes) <= k {
		reverse(sBytes)
		return string(sBytes)
	}

	left := 0
	right := min(2*k, len(sBytes))
	curSlice := sBytes[left:right]

	for left < len(s) {
		subK := min(k, right-left)
		reverse(curSlice[:subK])

		left += 2 * k
		if left >= len(sBytes) {
			break
		}

		right = min(left+2*k, len(sBytes))
		curSlice = sBytes[left:right]
	}

	return string(sBytes)
}

func reverse(b []byte) {
	lenb := len(b)
	for i := 0; i < lenb/2; i++ {
		b[i], b[lenb-1-i] = b[lenb-1-i], b[i]
	}
}
