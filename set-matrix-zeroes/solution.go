package setmatrixzeroes

func setZeroes(matrix [][]int) {
	zeroedRows := make(map[int]struct{})
	zeroedCols := make(map[int]struct{})
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				zeroedRows[row] = struct{}{}
				zeroedCols[col] = struct{}{}
			}
		}
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if _, ok := zeroedRows[row]; ok {
				matrix[row][col] = 0
			}
			if _, ok := zeroedCols[col]; ok {
				matrix[row][col] = 0
			}
		}
	}
}
