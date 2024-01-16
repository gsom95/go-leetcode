package rotatearray

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(n []int) {
	start, end := 0, len(n)-1
	for start < end {
		n[start], n[end] = n[end], n[start]
		start++
		end--
	}
}
