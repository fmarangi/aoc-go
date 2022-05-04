package day11

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestPowerLevel(t *testing.T) {
	utils.Assert(t, 4, Cell{3, 5}.Power(8))
	utils.Assert(t, -5, Cell{122, 79}.Power(57))
	utils.Assert(t, 0, Cell{217, 196}.Power(39))
	utils.Assert(t, 4, Cell{101, 153}.Power(71))
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, "20,34", Part1(input))
	if !testing.Short() {
		utils.Assert(t, "90,57,15", Part2(input))
	}
}
