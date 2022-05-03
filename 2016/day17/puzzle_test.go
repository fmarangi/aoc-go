package day17

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestShortest(t *testing.T) {
	utils.Assert(t, "DDRRRD", shortest("ihgpwlah"))
	utils.Assert(t, "DDUDRLRRUDRD", shortest("kglvqrro"))
	utils.Assert(t, "DRURDRUDDLLDLUURRDULRLDUUDDDRR", shortest("ulqzkmiv"))
}

func TestLongest(t *testing.T) {
	utils.Assert(t, 370, longest("ihgpwlah"))
	utils.Assert(t, 492, longest("kglvqrro"))
	utils.Assert(t, 830, longest("ulqzkmiv"))
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, "RDRDUDLRDR", Part1(input))
	utils.Assert(t, 386, Part2(input))
}
