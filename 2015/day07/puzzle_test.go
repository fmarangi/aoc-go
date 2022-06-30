package day07

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 46065, Part1(input))
	utils.Assert(t, 14134, Part2(input))
}
