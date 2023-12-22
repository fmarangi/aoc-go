package day11

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 9608724, Part1(input))
	utils.Assert(t, 904633799472, Part2(input))
}
