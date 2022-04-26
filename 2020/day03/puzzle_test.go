package day03

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 280, Part1(input))
	utils.Assert(t, 4355551200, Part2(input))
}
