package day05

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 6572, Part1(input))
	utils.Assert(t, 21466, Part2(input))
}
