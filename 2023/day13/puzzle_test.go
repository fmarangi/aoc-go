package day13

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 36015, Part1(input))
	utils.Assert(t, 35335, Part2(input))
}
