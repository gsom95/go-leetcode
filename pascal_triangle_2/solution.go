package pascaltriangle2

func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}

	if rowIndex == 0 {
		return []int{1}
	}
	result := row(1, rowIndex, []int{1, 1})

	return result
}

func row(curRowIndex, totalRows int, curRow []int) []int {
	newRow := make([]int, curRowIndex+1)
	newRow[0], newRow[curRowIndex] = 1, 1
	for i := 1; i < curRowIndex; i++ {
		newRow[i] = curRow[i-1] + curRow[i]
	}
	if curRowIndex == totalRows {
		return newRow
	}
	return row(curRowIndex+1, totalRows, newRow)
}
