package day17

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 600, Part1(input))
	utils.Assert(t, 31220910, Part2(input))
}
