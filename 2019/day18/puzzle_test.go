package day18

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 3646, Part1(input))
	utils.Assert(t, 1730, Part2(input))
}
