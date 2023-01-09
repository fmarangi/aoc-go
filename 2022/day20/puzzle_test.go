package day20

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 8372, Part1(input))
	if !testing.Short() {
		utils.Assert(t, 7865110481723, Part2(input))
	}
}
