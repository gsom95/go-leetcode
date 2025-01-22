package map_of_highest_peak

import (
	"container/heap"
)

type Cell struct {
	Row, Col, Height int
}

// PriorityQueue implementation
type PriorityQueue []Cell

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].Height < pq[j].Height }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Cell)) }
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[0 : n-1]
	return item
}

func highestPeak(isWater [][]int) [][]int {
	rows := len(isWater)
	cols := len(isWater[0])

	var (
		res     = make([][]int, rows)
		visited = make([][]bool, rows)
	)
	for i := range res {
		res[i] = make([]int, cols)
		visited[i] = make([]bool, cols)
	}

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	queue := make(PriorityQueue, 0)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if isWater[r][c] == 1 {
				queue = append(queue, Cell{
					Row:    r,
					Col:    c,
					Height: 0,
				})
			}
		}
	}

	for len(queue) > 0 {
		cell := heap.Pop(&queue).(Cell)
		if visited[cell.Row][cell.Col] {
			continue
		}

		res[cell.Row][cell.Col] = cell.Height
		visited[cell.Row][cell.Col] = true

		for _, dir := range dirs {
			neighbour := Cell{
				Row:    cell.Row + dir[0],
				Col:    cell.Col + dir[1],
				Height: cell.Height + 1,
			}

			if neighbour.Row >= 0 && neighbour.Row < rows && neighbour.Col >= 0 && neighbour.Col < cols &&
				!visited[neighbour.Row][neighbour.Col] {
				heap.Push(&queue, neighbour)
			}
		}
	}

	return res
}
