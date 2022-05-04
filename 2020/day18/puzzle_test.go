package day18

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestCalculate(t *testing.T) {
	utils.Assert(t, 71, Calculate("1 + 2 * 3 + 4 * 5 + 6"))
	utils.Assert(t, 51, Calculate("1 + (2 * 3) + (4 * (5 + 6))"))
	utils.Assert(t, 26, Calculate("2 * 3 + (4 * 5)"))
	utils.Assert(t, 437, Calculate("5 + (8 * 3 + 9 + 3 * 4 * 3)"))
	utils.Assert(t, 12240, Calculate("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"))
	utils.Assert(t, 13632, Calculate("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"))
}

func TestCalculateWithPrecedence(t *testing.T) {
	utils.Assert(t, 231, CalculateWithPrecedence("1 + 2 * 3 + 4 * 5 + 6"))
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 5374004645253, Part1(input))
	utils.Assert(t, 88782789402798, Part2(input))
}
