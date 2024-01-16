package reshapethematrix

func matrixReshape(mat [][]int, r int, c int) [][]int {
	mRows := len(mat)
	mCols := len(mat[0])
	if r*c != mRows*mCols {
		return mat
	}

	res := make([][]int, r)
	matRow := 0
	matCol := 0
	for rowIndex := range res {
		res[rowIndex] = make([]int, c)
		for resCol := range res[rowIndex] {
			res[rowIndex][resCol] = mat[matRow][matCol]
			matCol++
			if matCol == mCols {
				matCol = 0
				matRow++
			}
		}
	}

	return res
}
