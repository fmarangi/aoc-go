package day11

import (
	"strings"
)

func toInt(energy rune) int {
	return int(energy) - 48 // int('0')
}

func parseInput(input string) (octopuses []int, lineLength int) {
	for _, o := range input {
		if o != '\n' {
			octopuses = append(octopuses, toInt(o))
		}
	}
	return octopuses, strings.Index(input, "\n")
}

func next(octopuses []int, lineLength int) func(int) []int {
	max := len(octopuses)
	return func(i int) (around []int) {
		rowPos := i % lineLength
		for j := -1; j < 2; j++ {
			for k := -1; k < 2; k++ {
				a := i + j + lineLength*k
				if a >= 0 && a < max && a != i && rowPos+j >= 0 && rowPos+j < lineLength {
					around = append(around, a)
				}
			}
		}
		return
	}
}

func step(octopuses []int, next func(int) []int) ([]int, int) {
	flashes := []int{}
	for i, m := 0, len(octopuses); i < m; i++ {
		octopuses[i] = (octopuses[i] + 1) % 10
		if octopuses[i] == 0 {
			flashes = append(flashes, i)
		}
	}

	f := len(flashes)
	for ; len(flashes) > 0; flashes = flashes[1:] {
		curr := flashes[0]
		for _, n := range next(curr) {
			if octopuses[n] == 0 {
				continue
			}

			octopuses[n] = (octopuses[n] + 1) % 10
			if octopuses[n] == 0 {
				flashes = append(flashes, n)
				f++
			}
		}
	}
	return octopuses, f
}

func Part1(input string) (flashes int) {
	octopuses, lineLength := parseInput(input)
	around := next(octopuses, lineLength)
	for i, n := 0, 0; i < 100; i++ {
		octopuses, n = step(octopuses, around)
		flashes += n
	}
	return
}

func Part2(input string) int {
	octopuses, lineLength := parseInput(input)
	around := next(octopuses, lineLength)
	for i, n := 0, 0; ; i++ {
		octopuses, n = step(octopuses, around)
		if n == 100 {
			return i + 1
		}
	}
}
