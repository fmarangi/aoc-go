package day24

type Grid [5][5]bool

func (grid Grid) NextState() (next Grid) {
	offsets := [...]struct{ X, Y int }{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	for i := 0; i < 25; i++ {
		bugs, row, col := 0, i/5, i%5

		for _, o := range offsets {
			r, c := row+o.Y, col+o.X
			if r >= 0 && c >= 0 && r < 5 && c < 5 && grid[r][c] {
				bugs++
			}
		}

		if bugs == 1 || (bugs == 2 && !grid[row][col]) {
			next[row][col] = true
		}
	}

	return
}

func (grid Grid) Row(n int) (bugs int) {
	for i := 0; i < 5; i++ {
		if grid[n][i] {
			bugs++
		}
	}
	return
}

func (grid Grid) Col(n int) (bugs int) {
	for i := 0; i < 5; i++ {
		if grid[i][n] {
			bugs++
		}
	}
	return
}

func (grid Grid) Biodiversity() (result int) {
	for i := 0; i < 25; i++ {
		if grid[i/5][i%5] {
			result += 1 << i
		}
	}
	return
}

func (grid Grid) Bugs() (bugs int) {
	for i := 0; i < 25; i++ {
		if grid[i/5][i%5] {
			bugs++
		}
	}
	return
}

func (grid Grid) String() (result string) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if grid[i][j] {
				result += "#"
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return
}
