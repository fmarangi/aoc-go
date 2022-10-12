package day14

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 23890, Part1(input))
	if !testing.Short() {
		utils.Assert(t, 22696, Part2(input))
	}
}
