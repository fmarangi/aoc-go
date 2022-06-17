package day16

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 25895, Part1(input))
	utils.Assert(t, 5865723727753, Part2(input))
}
