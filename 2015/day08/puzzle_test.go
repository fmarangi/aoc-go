package day08

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestDecode(t *testing.T) {
	utils.Assert(t, ``, Decode(`""`))
	utils.Assert(t, `abc`, Decode(`"abc"`))
	utils.Assert(t, `aaa"aaa`, Decode(`"aaa\"aaa"`))
	utils.Assert(t, `'`, Decode(`"\x27"`))
}

func TestEncode(t *testing.T) {
	utils.Assert(t, `"\"\""`, Encode(`""`))
	utils.Assert(t, `"\"abc\""`, Encode(`"abc"`))
	utils.Assert(t, `"\"aaa\\\"aaa\""`, Encode(`"aaa\"aaa"`))
	utils.Assert(t, `"\"\\x27\""`, Encode(`"\x27"`))
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 1333, Part1(input))
	utils.Assert(t, 2046, Part2(input))
}
