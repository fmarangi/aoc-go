package day04

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Scratchcard struct {
	Winning, Numbers []int
	Copies           int
}

func (c Scratchcard) Matching() (matching int) {
	winning := make(map[int]bool, len(c.Winning))
	for _, n := range c.Winning {
		winning[n] = true
	}

	for _, n := range c.Numbers {
		if winning[n] {
			matching++
		}
	}

	return
}

func (c Scratchcard) Points() (points int) {
	for i, m := 0, c.Matching(); i < m; i++ {
		if points == 0 {
			points = 1
		} else {
			points *= 2
		}
	}
	return
}

func Part1(input string) (points int) {
	for _, c := range parseInput(input) {
		points += c.Points()
	}
	return
}

func Part2(input string) (total int) {
	cards := parseInput(input)
	for i, n := 0, len(cards); i < n; i++ {
		card := cards[i]
		if m := card.Matching(); m > 0 {
			for j := i + 1; j <= i+m && j < n; j++ {
				cards[j].Copies += card.Copies
			}
		}
	}

	for _, c := range cards {
		total += c.Copies
	}

	return
}

func parseInput(input string) (cards []Scratchcard) {
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		nums := strings.Split(parts[1], " | ")
		cards = append(cards, Scratchcard{
			Winning: utils.Ints(strings.Fields(nums[0])),
			Numbers: utils.Ints(strings.Fields(nums[1])),
			Copies:  1,
		})
	}
	return
}
