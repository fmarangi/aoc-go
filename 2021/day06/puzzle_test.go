package day06

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 350917, Part1(input))
	utils.Assert(t, 1592918715629, Part2(input))
}
