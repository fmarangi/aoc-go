package day01

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 54990, Part1(input))
	utils.Assert(t, 54473, Part2(input))
}
