package day23

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 4068, Part1(input))
	if !testing.Short() {
		utils.Assert(t, 968, Part2(input))
	}
}
