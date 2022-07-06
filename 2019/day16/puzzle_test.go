package day16

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 27831665, Part1(input))
	utils.Assert(t, 36265589, Part2(input))
}
