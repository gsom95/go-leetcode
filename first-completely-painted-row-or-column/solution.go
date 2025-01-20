package firstcompletelypaintedroworcolumn

func firstCompleteIndex(arr []int, mat [][]int) int {
	lenRows := len(mat)
	lenCols := len(mat[0])

	rows := make([]int, lenRows)
	cols := make([]int, lenCols)
	valuesMap := make(map[int][2]int)

	for rowIndex, row := range mat {
		for colIndex, elem := range row {
			valuesMap[elem] = [2]int{rowIndex, colIndex}
		}
	}

	for i, elem := range arr {
		rowIndex, colIndex := valuesMap[elem][0], valuesMap[elem][1]
		rows[rowIndex]++
		cols[colIndex]++
		if rows[rowIndex] == lenCols || cols[colIndex] == lenRows {
			return i
		}
	}

	return -1
}
