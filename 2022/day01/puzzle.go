package day01

import (
	"sort"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func Part1(input string) int {
	return parseInput(input)[0]
}

func Part2(input string) int {
	return utils.Sum(parseInput(input)[:3])
}

func parseInput(input string) (elves []int) {
	for _, elf := range strings.Split(input, "\n\n") {
		snacks := utils.Ints(strings.Fields(elf))
		elves = append(elves, utils.Sum(snacks))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	return
}
