package day13

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestCaught(t *testing.T) {
	utils.Assert(t, true, Scanner{0, 3}.Caught(0))
	utils.Assert(t, false, Scanner{1, 2}.Caught(0))
	utils.Assert(t, false, Scanner{4, 4}.Caught(0))
	utils.Assert(t, true, Scanner{6, 4}.Caught(0))
}

func TestEscape(t *testing.T) {
	scanners := []Scanner{{0, 3}, {1, 2}, {4, 4}, {6, 4}}
	for i := 0; i < 10; i++ {
		utils.Assert(t, false, Escape(scanners, i))
	}
	utils.Assert(t, true, Escape(scanners, 10))
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 1904, Part1(input))
	utils.Assert(t, 3833504, Part2(input))
}
