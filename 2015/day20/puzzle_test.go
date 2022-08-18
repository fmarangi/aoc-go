package day20

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 665280, Part1(input))
	utils.Assert(t, 705600, Part2(input))
}
