package day04

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestRoomIsValid(t *testing.T) {
	utils.Assert(t, true, Room{"aaaaa-bbb-z-y-x", "abxyz", 123}.IsValid())
	utils.Assert(t, true, Room{"a-b-c-d-e-f-g-h", "abcde", 987}.IsValid())
	utils.Assert(t, true, Room{"not-a-real-room", "oarel", 404}.IsValid())
	utils.Assert(t, false, Room{"totally-real-room", "decoy", 200}.IsValid())
}

func TestDecrypt(t *testing.T) {
	utils.Assert(t, "very encrypted name", Room{"qzmt-zixmtkozy-ivhz", "", 343}.Decrypt())
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 173787, Part1(input))
	utils.Assert(t, 548, Part2(input))
}
