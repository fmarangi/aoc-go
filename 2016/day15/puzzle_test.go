package day15

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 16824, Part1(input))
	utils.Assert(t, 3543984, Part2(input))
}
