package minimum_operations_to_exceed_threshold_value_ii

import (
	"container/heap"
)

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var _ heap.Interface = &Heap{}

func minOperations(nums []int, k int) int {
	operations := 0

	h := make(Heap, len(nums))
	copy(h, nums)
	heap.Init(&h)

	for x := heap.Pop(&h).(int); h.Len() > 0 && x < k; x = heap.Pop(&h).(int) {
		y := heap.Pop(&h).(int)
		operations++
		newElement := min(x, y)*2 + max(x, y)
		heap.Push(&h, newElement)
	}

	return operations
}
