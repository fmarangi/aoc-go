package day18

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 4192, Part1(input))
	utils.Assert(t, 2520, Part2(input))
}
