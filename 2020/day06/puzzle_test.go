package day06

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 7027, Part1(input))
	utils.Assert(t, 3579, Part2(input))
}
