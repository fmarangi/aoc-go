package day03

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 101469, Part1(input))
	utils.Assert(t, 1067, Part2(input))
}
