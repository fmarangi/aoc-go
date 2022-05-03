package day18

type Traps []bool

func (t Traps) Safe() (safe int) {
	for _, c := range t {
		if c {
			safe++
		}
	}
	return
}

func (current Traps) next() Traps {
	num := len(current)
	traps := make(Traps, num)

	for i := range current {
		var value int
		for j := -1; j < 2; j++ {
			if p := i + j; p >= 0 && p < num && !current[p] {
				value += 1 << (1 - j)
			}
		}

		switch value {
		case 0, 2, 5, 7:
			traps[i] = true
		}
	}

	return traps
}

func Part1(input string) (safe int) {
	traps := parseInput(input)
	safe += traps.Safe()

	for i := 1; i < 40; i++ {
		traps = traps.next()
		safe += traps.Safe()
	}
	return
}

func Part2(input string) (safe int) {
	traps := parseInput(input)
	safe += traps.Safe()

	for i := 1; i < 400000; i++ {
		traps = traps.next()
		safe += traps.Safe()
	}
	return
}

func parseInput(input string) Traps {
	traps := make(Traps, len(input))
	for i, c := range input {
		traps[i] = c != '^'
	}
	return traps
}
