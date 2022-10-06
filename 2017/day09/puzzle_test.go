package day09

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestScore(t *testing.T) {
	utils.Assert(t, 1, Score("{}"))
	utils.Assert(t, 6, Score("{{{}}}"))
	utils.Assert(t, 5, Score("{{},{}}"))
	utils.Assert(t, 16, Score("{{{},{},{{}}}}"))
	utils.Assert(t, 1, Score("{<a>,<a>,<a>,<a>}"))
	utils.Assert(t, 9, Score("{{<ab>},{<ab>},{<ab>},{<ab>}}"))
	utils.Assert(t, 9, Score("{{<!!>},{<!!>},{<!!>},{<!!>}}"))
	utils.Assert(t, 3, Score("{{<a!>},{<a!>},{<a!>},{<ab>}}"))
}

func TestGarbage(t *testing.T) {
	utils.Assert(t, 0, Garbage("<>"))
	utils.Assert(t, 17, Garbage("<random characters>"))
	utils.Assert(t, 3, Garbage("<<<<>"))
	utils.Assert(t, 2, Garbage("<{!>}>"))
	utils.Assert(t, 0, Garbage("<!!>"))
	utils.Assert(t, 0, Garbage("<!!!>>"))
	utils.Assert(t, 10, Garbage("<{o\"i!a,<{i<a>"))
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 11347, Part1(input))
	utils.Assert(t, 5404, Part2(input))
}
