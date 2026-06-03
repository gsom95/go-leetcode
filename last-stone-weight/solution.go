package laststoneweight

import "container/heap"

type intHeap []int

// Len implements [heap.Interface].
func (h intHeap) Len() int {
	return len(h)
}

// Less implements [heap.Interface].
func (h intHeap) Less(i int, j int) bool {
	return h[i] > h[j]
}

// Pop implements [heap.Interface].
func (h *intHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Push implements [heap.Interface].
func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Swap implements [heap.Interface].
func (h intHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

var _ heap.Interface = (*intHeap)(nil)

func lastStoneWeight(stones []int) int {
	h := intHeap(stones)
	heap.Init(&h)

	for h.Len() > 1 {
		y := heap.Pop(&h).(int)
		x := heap.Pop(&h).(int)

		if x != y {
			heap.Push(&h, y-x)
		}
	}

	if h.Len() == 0 {
		return 0
	}

	return h[0]
}
