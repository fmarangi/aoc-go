package day02

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, "33444", Part1(input))
	utils.Assert(t, "446A6", Part2(input))
}
