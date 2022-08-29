package day22

type Pos struct{ X, Y int }

const (
	North int = iota
	East
	South
	West
)

const (
	Clean int = iota
	Weakened
	Infected
	Flagged
)

var offsets = map[int]Pos{
	North: {0, -1},
	East:  {1, 0},
	South: {0, 1},
	West:  {-1, 0},
}

func Part1(input string) (bursts int) {
	var (
		pos, grid = parseInput(input)
		dir       = North
	)

	for i := 0; i < 10000; i++ {
		if grid[pos] {
			dir = (dir + 1) % 4
		} else {
			dir = (dir + 3) % 4
		}

		grid[pos] = !grid[pos]
		if grid[pos] {
			bursts++
		}

		pos.X += offsets[dir].X
		pos.Y += offsets[dir].Y
	}

	return
}

func Part2(input string) (bursts int) {
	var (
		grid          = make(map[Pos]int)
		pos, infected = parseInput(input)
		dir           = North
	)

	for p, status := range infected {
		if status {
			grid[p] = Infected
		}
	}

	for i := 0; i < 10000000; i++ {
		switch grid[pos] {
		case Infected:
			dir = (dir + 1) % 4
		case Flagged:
			dir = (dir + 2) % 4
		case Clean:
			dir = (dir + 3) % 4
		}

		grid[pos] = (grid[pos] + 1) % 4
		if grid[pos] == Infected {
			bursts++
		}

		pos.X += offsets[dir].X
		pos.Y += offsets[dir].Y
	}

	return
}

func parseInput(input string) (Pos, map[Pos]bool) {
	var (
		grid = make(map[Pos]bool)
		x, y int
	)
	for _, c := range input {
		switch c {
		case '\n':
			x, y = 0, y+1
		default:
			grid[Pos{x, y}] = c == '#'
			x++
		}
	}
	return Pos{x / 2, y / 2}, grid
}
