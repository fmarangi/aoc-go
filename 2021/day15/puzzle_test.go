package day15

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 626, Part1(input))
	if !testing.Short() {
		utils.Assert(t, 2966, Part2(input))
	}
}
