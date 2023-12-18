package day09

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 1877825184, Part1(input))
	utils.Assert(t, 1108, Part2(input))
}
