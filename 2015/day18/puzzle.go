package day18

const Size = 100

type Grid [Size][Size]bool

func (grid Grid) On() (on int) {
	for _, y := range grid {
		for _, x := range y {
			if x {
				on++
			}
		}
	}
	return
}

func (grid Grid) Next() (next Grid) {
	for i, n := 0, Size*Size; i < n; i++ {
		x, y, c := i%Size, i/Size, 0
		for j := 0; j < 9; j++ {
			a, b := x+j/3-1, y+j%3-1
			if a == x && b == y {
				continue
			}

			if a >= 0 && b >= 0 && a < Size && b < Size && grid[b][a] {
				c++
			}
		}

		next[y][x] = c == 3 || (c == 2 && grid[y][x])
	}
	return
}

func (grid Grid) Stuck() Grid {
	grid[0][0] = true
	grid[0][Size-1] = true
	grid[Size-1][0] = true
	grid[Size-1][Size-1] = true
	return grid
}

func Part1(input string) int {
	grid := parseInput(input)
	for i := 0; i < 100; i++ {
		grid = grid.Next()
	}
	return grid.On()
}

func Part2(input string) int {
	grid := parseInput(input).Stuck()
	for i := 0; i < 100; i++ {
		grid = grid.Next().Stuck()
	}
	return grid.On()
}

func parseInput(input string) (grid Grid) {
	var x, y int
	for _, c := range input {
		if c == '\n' {
			x, y = 0, y+1
		} else {
			grid[y][x] = c == '#'
			x++
		}
	}
	return
}
