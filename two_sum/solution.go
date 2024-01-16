package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, n := range nums {
		left := target - n
		if index, found := m[target-left]; found {
			return []int{index, i}
		}
		m[left] = i
	}

	return nil
}
