package day01

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 3381405, Part1(input))
	utils.Assert(t, 5069241, Part2(input))
}
