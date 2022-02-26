package day17

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 7750, Part1(input))
	utils.Assert(t, 4120, Part2(input))
}
