package day18

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 1961, Part1(input))
	utils.Assert(t, 20000795, Part2(input))
}
