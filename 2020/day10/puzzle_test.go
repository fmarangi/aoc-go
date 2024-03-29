package day10

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 2400, Part1(input))
	utils.Assert(t, 338510590509056, Part2(input))
}
