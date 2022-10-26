package day23

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 8281, Part1(input))
	utils.Assert(t, 911, Part2(input))
}
