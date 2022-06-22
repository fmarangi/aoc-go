package day10

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSequence(t *testing.T) {
	utils.Assert(t, "11", Sequence("1"))
	utils.Assert(t, "21", Sequence("11"))
	utils.Assert(t, "1211", Sequence("21"))
	utils.Assert(t, "111221", Sequence("1211"))
	utils.Assert(t, "312211", Sequence("111221"))
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 252594, Part1(input))
	utils.Assert(t, 3579328, Part2(input))
}
