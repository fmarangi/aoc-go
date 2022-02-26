package day07

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 335271, Part1(input))
	utils.Assert(t, 95851339, Part2(input))
}
