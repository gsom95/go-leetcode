package zeroarraytransformationiii

import (
	"cmp"
	"container/heap"
	"slices"
)

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxRemoval(nums []int, queries [][]int) int {
	slices.SortStableFunc(queries, func(i, j []int) int {
		return cmp.Compare(i[0], j[0])
	})

	pq := &Heap{}
	heap.Init(pq)
	deltaArray := make([]int, len(nums)+1)
	operations := 0

	for i, j := 0, 0; i < len(nums); i++ {
		operations += deltaArray[i]
		for j < len(queries) && queries[j][0] == i {
			heap.Push(pq, queries[j][1])
			j++
		}
		for operations < nums[i] && pq.Len() > 0 && (*pq)[0] >= i {
			operations += 1
			deltaArray[heap.Pop(pq).(int)+1] -= 1
		}
		if operations < nums[i] {
			return -1
		}
	}
	return pq.Len()
}
