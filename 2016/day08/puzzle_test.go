package day08

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 121, Part1(input))
	utils.Assert(t, "RURUCEOEIL", Part2(input))
}
