package day10

import "strings"

const (
	North int = iota
	East
	South
	West
)

var next = [4]map[byte]int{
	{'|': North, 'F': East, '7': West},
	{'-': East, 'J': North, '7': South},
	{'|': South, 'L': East, 'J': West},
	{'-': West, 'L': North, 'F': South},
}

func Part1(input string) (steps int) {
	var (
		start   = strings.Index(input, "S")
		width   = strings.Index(input, "\n") + 1
		offsets = [4]int{-width, 1, width, -1}
		max     = len(input)
		what    byte
		current int
		dir     int
	)

	for d, o := range offsets {
		current = start + o
		if current < 0 || current >= max {
			continue
		}

		what = input[current]
		if _, ok := next[d][what]; ok {
			dir = d
			break
		}
	}

	for steps = 1; what != 'S'; steps++ {
		current += offsets[dir]
		what = input[current]
		dir = next[dir][what]
	}

	return steps / 2
}

func Part2(input string) (summary int) {
	return
}
