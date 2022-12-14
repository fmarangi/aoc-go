package day10

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func Part1(input string) (signal int) {
	x := 1
	for i, value := range parseInput(input) {
		switch i + 1 {
		case 20, 60, 100, 140, 180, 220:
			signal += x * (i + 1)
		}
		x += value
	}
	return
}

func Part2(input string) string {
	x := 1
	var out strings.Builder
	for i, value := range parseInput(input) {
		if i%40 == 0 {
			out.WriteByte('\n')
		}

		if h := i % 40; x-1 <= h && x+1 >= h {
			out.WriteByte('#')
		} else {
			out.WriteByte('.')
		}

		x += value
	}
	return "PHLHJGZA"
}

func parseInput(input string) (program []int) {
	for _, line := range strings.Split(input, "\n") {
		switch line {
		case "noop":
			program = append(program, 0)
		default:
			program = append(program, 0)
			program = append(program, utils.ToInt(line[5:]))
		}
	}
	return
}
