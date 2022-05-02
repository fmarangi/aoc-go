package day19

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 1841611, Part1(input))
	utils.Assert(t, 1423634, Part2(input))
}
