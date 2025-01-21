package grid_game

import (
	"math"
)

func gridGame(grid [][]int) int64 {
	var (
		cols     = len(grid[0])
		firstSum int64
	)
	for _, v := range grid[0] {
		firstSum += int64(v)
	}

	var (
		secondSum int64
		minSum    int64 = math.MaxInt64
	)

	for turnIndex := 0; turnIndex < cols; turnIndex++ {
		firstSum -= int64(grid[0][turnIndex])
		minSum = min(minSum, max(firstSum, secondSum))
		secondSum += int64(grid[1][turnIndex])
	}

	return minSum
}
