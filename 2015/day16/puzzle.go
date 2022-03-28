package day16

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Sue map[string]int

func Facts() Sue {
	return Sue{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
}

func (a Sue) Match(b Sue) bool {
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func (a Sue) MatchWithRanges(b Sue) bool {
	for k, v := range a {
		switch k {
		case "cats", "trees":
			if b[k] >= v {
				return false
			}
		case "pomeranians", "goldfish":
			if b[k] <= v {
				return false
			}
		default:
			if b[k] != v {
				return false
			}
		}
	}
	return true
}

func Part1(input string) int {
	aunts, rightAunt := parseInput(input), Facts()
	for i, n := 0, len(aunts); i < n; i++ {
		if aunts[i].Match(rightAunt) {
			return i + 1
		}
	}
	return -1
}

func Part2(input string) int {
	aunts, rightAunt := parseInput(input), Facts()
	for i, n := 0, len(aunts); i < n; i++ {
		if aunts[i].MatchWithRanges(rightAunt) {
			return i + 1
		}
	}
	return -1
}

func parseInput(input string) (aunts []Sue) {
	for _, row := range strings.Split(input, "\n") {
		parts, sue := strings.SplitN(row, ": ", 2), Sue{}
		for _, thing := range strings.Split(parts[1], ", ") {
			data := strings.Split(thing, ": ")
			sue[data[0]] = utils.ToInt(data[1])
		}
		aunts = append(aunts, sue)
	}
	return
}
