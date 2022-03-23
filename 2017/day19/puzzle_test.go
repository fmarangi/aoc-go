package day19

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, "LIWQYKMRP", Part1(input))
	utils.Assert(t, 16764, Part2(input))
}
