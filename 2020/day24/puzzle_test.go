package day24

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 354, Part1(input))
	utils.Assert(t, 3608, Part2(input))
}
