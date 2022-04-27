package day09

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 25918798, Part1(input))
	utils.Assert(t, 3340942, Part2(input))
}
