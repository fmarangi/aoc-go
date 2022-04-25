package day05

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	if !testing.Short() {
		input := utils.ReadInput("input.txt")
		utils.Assert(t, "f77a0e6e", Part1(input))
		utils.Assert(t, "999828ec", Part2(input))
	}
}
