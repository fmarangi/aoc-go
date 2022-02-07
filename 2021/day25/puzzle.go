package day25

import (
	"strings"
)

type AreaMap struct {
	locations  []rune
	lineLength int
}

func (a AreaMap) String() string {
	return string(a.locations)
}

func (a *AreaMap) Step() (moved bool) {
	max := len(a.locations)
	newMap := make([]rune, max)

	// Move east-facing
	for i := 0; i < max; i++ {
		if a.locations[i] == '>' {
			target := i + 1
			if target == max || a.locations[target] == '\n' {
				target -= a.lineLength - 1
			}
			if a.locations[target] == '.' {
				newMap[i] = '.'
				newMap[target] = '>'
				moved = true
				continue
			}
		}

		if newMap[i] != '>' {
			newMap[i] = a.locations[i]
		}
	}
	a.locations = newMap

	// Move south-facing
	newMap = make([]rune, max)
	for i := 0; i < max; i++ {
		if a.locations[i] == 'v' {
			target := (i + a.lineLength) % (max + 1)
			if a.locations[target] == '.' {
				newMap[i] = '.'
				newMap[target] = 'v'
				moved = true
				continue
			}
		}

		if newMap[i] != 'v' {
			newMap[i] = a.locations[i]
		}
	}
	a.locations = newMap

	return
}

func parseInput(input string) AreaMap {
	line := strings.Index(input, "\n") + 1
	return AreaMap{[]rune(input), line}
}

func Part1(input string) (steps int) {
	a := parseInput(input)
	for a.Step() {
		steps++
	}
	return steps + 1
}
