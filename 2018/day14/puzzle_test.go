package day14

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestRecipes(t *testing.T) {
	utils.Assert(t, 5158916779, Scores(9))
	utils.Assert(t, 124515891, Scores(5))
	utils.Assert(t, 9251071085, Scores(18))
	utils.Assert(t, 5941429882, Scores(2018))
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 7861362411, Part1(input))
	if !testing.Short() {
		utils.Assert(t, 20203532, Part2(input))
	}
}
