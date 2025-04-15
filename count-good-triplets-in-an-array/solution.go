package countgoodtripletsinanarray

type FenwickTree struct {
	tree []int
}

func NewFenwickTree(size int) *FenwickTree {
	return &FenwickTree{tree: make([]int, size+1)} // indexing is from 1
}

func (t *FenwickTree) update(index, value int) {
	index += 1
	for index < len(t.tree) {
		t.tree[index] += value
		index += index & -index
	}
}

func (t *FenwickTree) sum(index int) int {
	if index < 0 {
		return 0
	}

	index++
	result := 0
	for index > 0 {
		result += t.tree[index]
		index -= index & -index
	}
	return result
}

func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	pos1 := make(map[int]int, n)

	for i := range nums1 {
		pos1[nums1[i]] = i

	}

	mapped := make([]int, n)
	for i := 0; i < n; i++ {
		mapped[i] = pos1[nums2[i]]
	}

	smaller := make([]int, n)
	bitLeft := NewFenwickTree(n)
	for i := 0; i < n; i++ {
		smaller[i] = bitLeft.sum(mapped[i] - 1)
		bitLeft.update(mapped[i], 1)
	}
	larger := make([]int, n)
	bitRight := NewFenwickTree(n)
	for i := n - 1; i >= 0; i-- {
		larger[i] = bitRight.sum(n-1) - bitRight.sum(mapped[i])
		bitRight.update(mapped[i], 1)
	}

	res := 0
	for i := range n {
		res += smaller[i] * larger[i]
	}

	return int64(res)
}
