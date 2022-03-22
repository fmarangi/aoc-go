package day06

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 11137, Part1(input))
	utils.Assert(t, 1037, Part2(input))
}
