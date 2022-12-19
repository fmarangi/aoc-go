package day11

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 58322, Part1(input))
	utils.Assert(t, 13937702909, Part2(input))
}
