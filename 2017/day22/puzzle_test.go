package day22

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 5196, Part1(input))
	utils.Assert(t, 2511633, Part2(input))
}
