package day25

import "strings"

type Pins [5]uint8

func (lock Pins) Fits(key Pins) bool {
	for i, n := range lock {
		if n+key[i] > 5 {
			return false
		}
	}
	return true
}

func Part1(input string) (fit int) {
	locks, keys := parseInput(input)
	for _, lock := range locks {
		for _, key := range keys {
			if lock.Fits(key) {
				fit++
			}
		}
	}
	return
}

func parseInput(input string) (keys, locks []Pins) {
	for _, schema := range strings.Split(input, "\n\n") {
		var pins Pins
		for _, row := range strings.Fields(schema)[1:6] {
			for i, c := range row {
				if c == '#' {
					pins[i]++
				}
			}
		}

		if schema[0] == '#' {
			locks = append(locks, pins)
		} else {
			keys = append(keys, pins)
		}
	}
	return
}
