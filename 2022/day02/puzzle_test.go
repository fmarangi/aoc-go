package day02

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 12156, Part1(input))
	utils.Assert(t, 10835, Part2(input))
}
