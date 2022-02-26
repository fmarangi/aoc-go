package day20

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 5395, Part1(input))
	// utils.Assert(t, 17584, Part2(input))
}
