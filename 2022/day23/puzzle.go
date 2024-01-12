package day23

type Position struct{ X, Y int }

func (a Position) Move(b Position) Position {
	return Position{a.X + b.X, a.Y + b.Y}
}

func Part1(input string) int {
	var grove = parseInput(input)
	for i := 0; i < 10; i++ {
		grove, _ = grove.Next(i)
	}
	return grove.EmptyTiles()
}

func Part2(input string) (round int) {
	grove := parseInput(input)
	for moved := true; moved; round++ {
		grove, moved = grove.Next(round)
	}
	return
}

func parseInput(input string) Grove {
	var (
		x, y  int
		grove = make(Grove)
	)

	for _, pos := range input {
		switch pos {
		case '\n':
			x, y = 0, y+1
			continue
		case '#':
			grove[Position{x, y}] = true
		}
		x++
	}

	return grove
}
