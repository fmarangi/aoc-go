package day09

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestScore(t *testing.T) {
	utils.Assert(t, Score("{}"), 1)
	utils.Assert(t, Score("{{{}}}"), 6)
	utils.Assert(t, Score("{{},{}}"), 5)
	utils.Assert(t, Score("{{{},{},{{}}}}"), 16)
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 11347, Part1(input))
	utils.Assert(t, 5404, Part2(input))
}
