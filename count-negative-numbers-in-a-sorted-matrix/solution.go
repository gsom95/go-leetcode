package count_negative_numbers_in_a_sorted_matrix

func countNegatives(grid [][]int) int {
	res := 0

	rowLen := len(grid)
	curRow := 0

	colLen := len(grid[0])
	curCol := colLen - 1

	for curCol >= 0 && curRow < rowLen {
		if grid[curRow][curCol] >= 0 {
			curRow++
		} else {
			res += rowLen - curRow
			curCol--
		}
	}

	return res
}
