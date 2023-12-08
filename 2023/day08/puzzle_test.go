package day08

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 11309, Part1(input))
	utils.Assert(t, 13740108158591, Part2(input))
}
