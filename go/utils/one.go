package utils

// t: O(n), s: O(n)
func findDiagonalOrder(mat [][]int) []int {
	result := make([]int, len(mat)*len(mat))

	getCoordinates := coordinatesClosure(len(mat))
	var row, col int

	for i := 0; i < len(result); i++ {
		row, col = getCoordinates()
		result[i] = mat[row][col]
	}

	return result
}

func updateCoordinates(increment bool, row, col, limit int) (bool, int, int) {
	if increment {
		row--
		col = max(col, col+1, limit)
		if row == -1 {
			row = 0
			increment = false
		}
	} else {
		row = max(row, row+1, limit)
		col--
		if col == -1 {
			col = 0
			increment = true
		}
	}

	return increment, row, col
}

func max(n, n1, limit int) int {
	result := n
	if n1 > n {
		result = n1
	}
	if result >= limit {
		result--
	}
	return result
}

func coordinatesClosure(limit int) func() (int, int) {
	goUp, row, col := true, 1, -1

	f := func() (int, int) {
		if goUp {
			if row == 0 || col+1 == limit {
				goUp = false
				if col+1 == limit {
					row++
				} else {
					col++
				}
			} else {
				row--
				col++
			}
		} else {
			if col == 0 || row+1 == limit {
				goUp = true
				if row+1 == limit {
					col++
				} else {
					row++
				}
			} else {
				row++
				col--
			}
		}

		return row, col
	}

	return f
}
