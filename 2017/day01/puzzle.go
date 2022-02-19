package day01

import "github.com/fmarangi/aoc-go/utils"

func Part1(input string) int {
	return CountMatches(input, 1)
}

func Part2(input string) int {
	return CountMatches(input, len(input)/2)
}

func CountMatches(input string, offset int) (match int) {
	n := len(input)
	for i, c := range input {
		if c == rune(input[(i+offset)%n]) {
			match += utils.ToInt(string(c))
		}
	}
	return
}
