package day15

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 505379, Part1(input))
	utils.Assert(t, 263211, Part2(input))
}
