package day06

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, "ygjzvzib", Part1(input))
	utils.Assert(t, "pdesmnoz", Part2(input))
}
