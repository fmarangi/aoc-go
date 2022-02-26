package day24

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 94992992796199, Part1(input))
	utils.Assert(t, 11931881141161, Part2(input))
}
