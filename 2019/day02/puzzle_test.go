package day02

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 3409710, Part1(input))
	utils.Assert(t, 7912, Part2(input))
}
