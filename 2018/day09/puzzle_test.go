package day09

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestGame(t *testing.T) {
	utils.Assert(t, 8317, Game(1618, 10))
	utils.Assert(t, 146373, Game(7999, 13))
	utils.Assert(t, 2764, Game(1104, 17))
	utils.Assert(t, 54718, Game(6111, 21))
	utils.Assert(t, 37305, Game(5807, 30))
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 429943, Part1(input))
	utils.Assert(t, 3615691746, Part2(input))
}
