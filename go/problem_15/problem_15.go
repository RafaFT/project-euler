/*
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down,
there are exactly 6 routes to the bottom right corner.

How many such routes are there through a 20×20 grid?
*/

package problem15

type matrix struct {
	maxRow    int
	maxColumn int
}

func newMatrix(gridSize int) matrix {
	num := gridSize + 1

	return matrix{
		maxRow:    num,
		maxColumn: num,
	}
}

func (m *matrix) calcPaths(row, column int) int {
	if row == m.maxRow || column == m.maxColumn {
		return 1
	}

	if row > m.maxRow || column > m.maxColumn {
		return 0
	}

	return m.calcPaths(row, column+1) + m.calcPaths(row+1, column)
}

// An attempt of divide and conquer algorithm, which gives the correct
// answer, but takes too long. Grid od 20 takes around 10 minutes on my PC.
func problem15() int {
	m := newMatrix(20)

	return m.calcPaths(1, 1)
}
