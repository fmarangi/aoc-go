package day01

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func parseInput(input string) []int {
	return utils.Ints(strings.Fields(input))
}

func Part1(input string) int {
	return utils.Sum(parseInput(input))
}

func Part2(input string) (frequency int) {
	frequencies, nums := map[int]bool{}, parseInput(input)
	for i, max := 0, len(nums); !frequencies[frequency]; i++ {
		frequencies[frequency] = true
		frequency += nums[i%max]
	}
	return
}
