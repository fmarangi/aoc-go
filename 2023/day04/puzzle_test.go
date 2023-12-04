package day04

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 23750, Part1(input))
	utils.Assert(t, 13261850, Part2(input))
}
