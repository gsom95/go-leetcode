package neighboringbitwisexor

func doesValidArrayExist(derived []int) bool {
	var res int

	for _, bit := range derived {
		res ^= bit
	}

	return res == 0
}
