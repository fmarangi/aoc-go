package day14

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 6631883285184, Part1(input))
	utils.Assert(t, 3161838538691, Part2(input))
}
