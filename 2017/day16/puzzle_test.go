package day16

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, "kpfonjglcibaedhm", Part1(input))
	utils.Assert(t, "odiabmplhfgjcekn", Part2(input))
}
