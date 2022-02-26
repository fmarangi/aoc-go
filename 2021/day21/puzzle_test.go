package day21

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 888735, Part1(input))
	utils.Assert(t, 647608359455719, Part2(input))
}
