package day14

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 768, Part1(input))
	utils.Assert(t, 26686, Part2(input))
}
