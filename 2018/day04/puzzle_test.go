package day04

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 151754, Part1(input))
	utils.Assert(t, 19896, Part2(input))
}
