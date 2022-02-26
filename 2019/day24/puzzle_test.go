package day24

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 28903899, Part1(input))
	utils.Assert(t, 1896, Part2(input))
}
