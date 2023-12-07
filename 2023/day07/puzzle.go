package day07

import (
	"sort"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func Part1(input string) (winnings int) {
	hands := parseInput(input)
	sort.Slice(hands, func(i, j int) bool {
		a, b := hands[i], hands[j]
		if a.Type == b.Type {
			for i := 0; i < 5; i++ {
				va, vb := cardValues[a.Cards[i]], cardValues[b.Cards[i]]
				if va == vb {
					continue
				}
				return va < vb
			}
		}
		return a.Type < b.Type
	})

	for i, h := range hands {
		winnings += h.Bid * (i + 1)
	}

	return
}

func Part2(input string) (winnings int) {
	hands := parseInputPart2(input)
	sort.Slice(hands, func(i, j int) bool {
		a, b := hands[i], hands[j]
		if a.Type == b.Type {
			for i := 0; i < 5; i++ {
				va, vb := cardValuesWithJoker[a.Cards[i]], cardValuesWithJoker[b.Cards[i]]
				if va == vb {
					continue
				}
				return va < vb
			}
		}
		return a.Type < b.Type
	})

	for i, h := range hands {
		winnings += h.Bid * (i + 1)
	}

	return
}

func parseInput(input string) (hands []Hand) {
	data := strings.Fields(input)
	for i, n := 0, len(data); i < n; i += 2 {
		hands = append(hands, Hand{data[i], utils.ToInt(data[i+1]), HandType(data[i])})
	}
	return
}

func parseInputPart2(input string) (hands []Hand) {
	data := strings.Fields(input)
	for i, n := 0, len(data); i < n; i += 2 {
		hands = append(hands, Hand{data[i], utils.ToInt(data[i+1]), HandTypeWithJoker(data[i])})
	}
	return
}
