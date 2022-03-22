package day05

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func Part1(input string) (steps int) {
	offsets := parseInput(input)
	for i, n := 0, len(offsets); i < n; steps++ {
		i, offsets[i] = i+offsets[i], offsets[i]+1
	}
	return
}

func Part2(input string) (steps int) {
	offsets := parseInput(input)
	for i, n := 0, len(offsets); i < n; steps++ {
		if offsets[i] < 3 {
			i, offsets[i] = i+offsets[i], offsets[i]+1
		} else {
			i, offsets[i] = i+offsets[i], offsets[i]-1
		}
	}
	return
}

func parseInput(input string) []int {
	return utils.Ints(strings.Fields(input))
}
