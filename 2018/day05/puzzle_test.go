package day05

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 10180, Part1(input))
	if !testing.Short() {
		utils.Assert(t, 5668, Part2(input))
	}
}
