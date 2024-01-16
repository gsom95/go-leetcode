package validsudoku

func isValidSudoku(board [][]byte) bool {
	const size = 9

	for row := 0; row < size; row++ {
		seenValues := make(map[byte]struct{}, size)
		for _, val := range board[row] {
			if val == '.' {
				continue
			}
			if _, seen := seenValues[val]; seen {
				return false
			}
			seenValues[val] = struct{}{}
		}
	}
	for col := 0; col < size; col++ {
		seenValues := make(map[byte]struct{}, size)
		for row := 0; row < size; row++ {
			val := board[row][col]
			if val == '.' {
				continue
			}
			if _, seen := seenValues[val]; seen {
				return false
			}
			seenValues[val] = struct{}{}
		}
	}

	for row := 0; row < size; row += 3 {
		for col := 0; col < size; col += 3 {
			seenValues := make(map[byte]struct{}, size)
			for gridRow := row; gridRow < row+3; gridRow++ {
				for gridCol := col; gridCol < col+3; gridCol++ {
					val := board[gridRow][gridCol]
					if val == '.' {
						continue
					}
					if _, seen := seenValues[val]; seen {
						return false
					}
					seenValues[val] = struct{}{}
				}
			}
		}
	}

	return true
}
