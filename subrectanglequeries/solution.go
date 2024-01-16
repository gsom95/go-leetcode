package subrectanglequeries

type SubrectangleQueries struct {
	matrix [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	m := SubrectangleQueries{rectangle}
	for i := range rectangle {
		m.matrix = append(m.matrix, make([]int, len(rectangle[i])))
		copy(m.matrix[i], rectangle[i])
	}

	return m
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for row := row1; row <= row2; row++ {
		for col := col1; col <= col2; col++ {
			this.matrix[row][col] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.matrix[row][col]
}
