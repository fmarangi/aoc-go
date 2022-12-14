package day10

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 15360, Part1(input))
	utils.Assert(t, "PHLHJGZA", Part2(input))
}
