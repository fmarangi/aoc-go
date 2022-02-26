package day12

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 3497, Part1(input))
	utils.Assert(t, 93686, Part2(input))
}
