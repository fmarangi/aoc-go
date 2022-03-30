package day11

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestNext(t *testing.T) {
	utils.Assert(t, Next("zzzzzzzz"), "aaaaaaaa")
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, "vzbxxyzz", Part1(input))
	utils.Assert(t, "vzcaabcc", Part2(input))
}
