package day12

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 8287, Part1(input))
	utils.Assert(t, 528250271633772, Part2(input))
}
