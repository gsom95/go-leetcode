package diagonaltraverse

func findDiagonalOrder(mat [][]int) []int {
	var (
		rows           = len(mat)
		cols           = len(mat[0])
		matElems       = rows * cols
		res            = make([]int, matElems)
		resIndex       = 0
		curRow, curCol = 0, 0
		upDirection    = true
	)

	for resIndex < matElems {
		res[resIndex] = mat[curRow][curCol]
		resIndex++

		if upDirection {
			if curCol == cols-1 {
				curRow = curRow + 1
				upDirection = false
			} else if curRow == 0 {
				curCol++
				upDirection = false
			} else {
				curRow--
				curCol++
			}
			continue
		}

		if curRow == rows-1 {
			curCol++
			upDirection = true
		} else if curCol == 0 {
			curRow++
			upDirection = true
		} else {
			curRow++
			curCol--
		}
	}

	return res
}
