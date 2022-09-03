package day03

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 438, Part1(input))
	utils.Assert(t, 266330, Part2(input))
}
