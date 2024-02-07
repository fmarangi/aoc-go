package day24

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	if !testing.Short() {
		utils.Assert(t, 11846773891, Part1(input))
		utils.Assert(t, 80393059, Part2(input))
	}
}
