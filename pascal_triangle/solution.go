package pascaltriangle

func row(curRowIndex int, curRow []int) []int {
	newRow := make([]int, curRowIndex+1)
	newRow[0], newRow[curRowIndex] = 1, 1
	for i := 1; i < curRowIndex; i++ {
		newRow[i] = curRow[i-1] + curRow[i]
	}
	return newRow
}

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	res := make([][]int, numRows)
	res[0] = []int{1}
	res[1] = []int{1, 1}
	for r := 2; r < numRows; r++ {
		res[r] = row(r, res[r-1])
	}

	return res
}
