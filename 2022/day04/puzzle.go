package day04

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Section struct{ From, To int }

func (a Section) Contains(b Section) bool {
	return a.From <= b.From && a.To >= b.To
}

func (a Section) Overlaps(b Section) bool {
	return (a.From <= b.From && a.To >= b.From) || (b.From <= a.From && b.To >= a.From)
}

func Part1(input string) (count int) {
	for _, pair := range parseInput(input) {
		if pair[0].Contains(pair[1]) || pair[1].Contains(pair[0]) {
			count++
		}
	}
	return
}

func Part2(input string) (count int) {
	for _, pair := range parseInput(input) {
		if pair[0].Overlaps(pair[1]) {
			count++
		}
	}
	return
}

func parseInput(input string) (pairs [][2]Section) {
	for _, row := range strings.Fields(input) {
		pair := strings.Split(row, ",")
		pairs = append(pairs, [2]Section{
			parseSection(pair[0]),
			parseSection(pair[1]),
		})
	}
	return
}

func parseSection(input string) Section {
	data := utils.Ints(strings.Split(input, "-"))
	return Section{data[0], data[1]}
}
