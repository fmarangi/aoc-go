package day15

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	if testing.Short() {
		t.Skip("calculating both parts takes up to 30s")
	}

	input := utils.ReadInput("input.txt")
	utils.Assert(t, 638, Part1(input))
	utils.Assert(t, 343, Part2(input))
}
