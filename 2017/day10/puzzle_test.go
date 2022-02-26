package day10

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 8536, Part1(input))
	utils.Assert(t, "aff593797989d665349efe11bb4fd99b", Part2(input))
}
