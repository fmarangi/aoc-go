package day07

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestHandType(t *testing.T) {
	utils.Assert(t, OnePair, HandType("32T3K"))
	utils.Assert(t, ThreeOfAKind, HandType("T55J5"))
	utils.Assert(t, TwoPair, HandType("KK677"))
	utils.Assert(t, TwoPair, HandType("KTJJT"))
	utils.Assert(t, ThreeOfAKind, HandType("QQQJA"))
}

func TestHandTypeWithJoker(t *testing.T) {
	utils.Assert(t, OnePair, HandTypeWithJoker("32T3K"))
	utils.Assert(t, FourOfAKind, HandTypeWithJoker("T55J5"))
	utils.Assert(t, TwoPair, HandTypeWithJoker("KK677"))
	utils.Assert(t, FourOfAKind, HandTypeWithJoker("KTJJT"))
	utils.Assert(t, FourOfAKind, HandTypeWithJoker("QQQJA"))
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 253866470, Part1(input))
	utils.Assert(t, 254494947, Part2(input))
}
