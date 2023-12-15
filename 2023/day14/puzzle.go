package day14

func Part1(input string) int {
	return parseInput(input).Tilt(North).Load()
}

func Part2(input string) (summary int) {
	var (
		platform = parseInput(input)
		seen     = make(map[string]int)
		p        = platform.String()
	)

	// Find loop
	for seen[p] != 2 {
		seen[p]++
		platform = platform.Cycle()
		p = platform.String()
	}

	var head, size int
	for _, s := range seen {
		switch s {
		case 1:
			head++
		default:
			size++
		}
	}

	n := head + (1000000000-head)%size
	platform = parseInput(input)
	for i := 0; i < n; i++ {
		platform = platform.Cycle()
	}

	return platform.Load()
}

func parseInput(input string) (platform Platform) {
	platform = make(Platform)
	var x, y int
	for _, p := range input {
		switch p {
		case '\n':
			x, y = -1, y+1
		case '#':
			platform[Position{x, y}] = Wall
		case 'O':
			platform[Position{x, y}] = Rock
		}
		x++
	}
	return
}
