package search2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	cols := len(matrix[0])

	left := 0
	totalElements := len(matrix) * cols
	right := totalElements

	for left <= right {
		mid := (left + right) / 2

		if mid >= totalElements {
			return false
		}

		midVal := matrix[mid/cols][mid%cols]
		if target == midVal {
			return true
		}
		if target > midVal {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
