package day21

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestOperations(t *testing.T) {
	utils.Assert(t, "ebcda", Swap("abcde", 4, 0))
	utils.Assert(t, "bcdea", Rotate("abcde", -1))
	utils.Assert(t, "bdeac", Move("bcdea", 1, 4))
	utils.Assert(t, "abdec", Move("bdeac", 3, 0))
	utils.Assert(t, "ecabd", RotateLetter("abdec", "b"))
	utils.Assert(t, "decab", RotateLetter("ecabd", "d"))
	utils.Assert(t, "abcde", Reverse("edcba", 0, 4))
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, "gbhafcde", Part1(input))
	utils.Assert(t, "bcfaegdh", Part2(input))
}
