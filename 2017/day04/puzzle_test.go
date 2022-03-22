package day04

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestHasDuplicates(t *testing.T) {
	utils.Assert(t, hasDuplicates("aa bb cc dd ee"), false)
	utils.Assert(t, hasDuplicates("aa bb cc dd aa"), true)
	utils.Assert(t, hasDuplicates("aa bb cc dd aaa"), false)
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 455, Part1(input))
	utils.Assert(t, 186, Part2(input))
}
