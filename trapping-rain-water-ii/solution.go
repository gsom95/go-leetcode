package trappingrainwaterii

import "container/heap"

// Cell represents a position in the heightMap
type Cell struct {
	x, y, height int
}

// PriorityQueue implementation
type PriorityQueue []Cell

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].height < pq[j].height }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Cell)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) <= 2 || len(heightMap[0]) <= 2 {
		return 0
	}

	m, n := len(heightMap), len(heightMap[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// Initialize priority queue with border cells
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Add border cells
	for i := 0; i < m; i++ {
		heap.Push(pq, Cell{i, 0, heightMap[i][0]})
		heap.Push(pq, Cell{i, n - 1, heightMap[i][n-1]})
		visited[i][0] = true
		visited[i][n-1] = true
	}
	for j := 1; j < n-1; j++ {
		heap.Push(pq, Cell{0, j, heightMap[0][j]})
		heap.Push(pq, Cell{m - 1, j, heightMap[m-1][j]})
		visited[0][j] = true
		visited[m-1][j] = true
	}

	// directions for adjacent cells
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	water := 0

	// Process cells from outside to inside
	for pq.Len() > 0 {
		cell := heap.Pop(pq).(Cell)

		// Check all adjacent cells
		for _, dir := range dirs {
			nx, ny := cell.x+dir[0], cell.y+dir[1]

			if nx >= 0 && nx < m && ny >= 0 && ny < n && !visited[nx][ny] {
				visited[nx][ny] = true

				// Add water if possible
				if heightMap[nx][ny] < cell.height {
					water += cell.height - heightMap[nx][ny]
					heap.Push(pq, Cell{nx, ny, cell.height})
				} else {
					heap.Push(pq, Cell{nx, ny, heightMap[nx][ny]})
				}
			}
		}
	}

	return water
}
