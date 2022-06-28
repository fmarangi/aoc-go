package day04

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 1864, Part1(input))
	utils.Assert(t, 1258, Part2(input))
}
