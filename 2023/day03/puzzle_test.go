package day03

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestPartNumbers(t *testing.T) {
	pn := partNumbers(utils.ReadInput("sample.txt"))
	utils.Assert(t, true, pn(0, 3) > -1)
	utils.Assert(t, true, pn(24, 2) > -1)
	utils.Assert(t, true, pn(28, 3) > -1)
	utils.Assert(t, true, pn(44, 3) > -1)
	utils.Assert(t, false, pn(5, 3) > -1)
	utils.Assert(t, false, pn(62, 2) > -1)
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 527364, Part1(input))
	utils.Assert(t, 79026871, Part2(input))
}
