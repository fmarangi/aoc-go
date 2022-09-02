package day06

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	if !testing.Short() {
		input := utils.ReadInput("input.txt")
		utils.Assert(t, 569999, Part1(input))
		utils.Assert(t, 17836115, Part2(input))
	}
}
