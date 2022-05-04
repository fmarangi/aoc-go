package day16

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, "10011010010010010", Part1(input))
	if !testing.Short() {
		utils.Assert(t, "10101011110100011", Part2(input))
	}
}
