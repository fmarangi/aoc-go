package day11

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestMove(t *testing.T) {
	utils.Assert(t, 3, Move("ne", "ne", "ne"))
	utils.Assert(t, 0, Move("ne", "ne", "sw", "sw"))
	utils.Assert(t, 2, Move("ne", "ne", "s", "s"))
	utils.Assert(t, 3, Move("se", "sw", "se", "sw", "sw"))
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 794, Part1(input))
	utils.Assert(t, 1524, Part2(input))
}
