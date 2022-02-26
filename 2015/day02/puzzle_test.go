package day02

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 1606483, Part1(input))
	utils.Assert(t, 3842356, Part2(input))
}
