package day22

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 602574, Part1(input))
	utils.Assert(t, 1288707160324706, Part2(input))
}
