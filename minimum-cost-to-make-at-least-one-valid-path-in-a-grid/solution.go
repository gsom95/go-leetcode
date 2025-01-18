package minimumcosttomakeatleastonevalidpathinagrid

import (
	"container/heap"
	"math"
)

type Cell struct {
	x, y, cost int
}

type MinHeap []Cell

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Cell))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	directions := []struct {
		dx, dy, dir int
	}{
		{0, 1, 1},  // right
		{0, -1, 2}, // left
		{1, 0, 3},  // down
		{-1, 0, 4}, // up
	}

	minCost := make([][]int, m)
	for i := range minCost {
		minCost[i] = make([]int, n)
		for j := range minCost[i] {
			minCost[i][j] = math.MaxInt // Initialize with a large number
		}
	}

	h := &MinHeap{}
	heap.Push(h, Cell{0, 0, 0})
	minCost[0][0] = 0

	for h.Len() > 0 {
		cell := heap.Pop(h).(Cell)
		x, y, cost := cell.x, cell.y, cell.cost

		if x == m-1 && y == n-1 {
			return cost
		}

		for _, d := range directions {
			nx, ny := x+d.dx, y+d.dy
			if nx >= 0 && nx < m && ny >= 0 && ny < n {
				newCost := cost
				if grid[x][y] != d.dir {
					newCost++
				}
				if newCost < minCost[nx][ny] {
					minCost[nx][ny] = newCost
					heap.Push(h, Cell{nx, ny, newCost})
				}
			}
		}
	}

	return -1 // This should never be reached if there is a valid path
}
