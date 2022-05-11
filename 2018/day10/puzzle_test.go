package day10

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, "RLEZNRAN", Part1(input))
	utils.Assert(t, 10240, Part2(input))
}
