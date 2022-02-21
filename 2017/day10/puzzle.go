package day10

import (
	"strings"

	"github.com/fmarangi/aoc-go/2017/knot"
	"github.com/fmarangi/aoc-go/utils"
)

func Part1(input string) int {
	circle, skip, pos := knot.Init(), 0, 0
	for _, l := range parseInput(input) {
		circle.Reverse(pos, l)
		pos += l + skip
		skip++
	}
	return int(circle[0]) * int(circle[1])
}

func Part2(input string) string {
	return knot.Hash(input)
}

func parseInput(input string) []int {
	return utils.Ints(strings.Split(input, ","))
}
