package day04

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 346386, Part1(input))
	if !testing.Short() {
		utils.Assert(t, 9958218, Part2(input))
	}
}
