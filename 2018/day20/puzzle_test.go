package day20

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 3314, Part1(input))
	utils.Assert(t, 8550, Part2(input))
}
