package day19

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 330, Part1(input))
	utils.Assert(t, 9634, Part2(input))
}
