package day02

import "strings"

func Part1(input string) (score int) {
	results := map[string]int{
		"A X": 4, "A Y": 8, "A Z": 3,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 7, "C Y": 2, "C Z": 6,
	}
	for _, round := range strings.Split(input, "\n") {
		score += results[round]
	}
	return
}

func Part2(input string) (score int) {
	results := map[string]int{
		"A X": 3, "A Y": 4, "A Z": 8,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 2, "C Y": 6, "C Z": 7,
	}
	for _, round := range strings.Split(input, "\n") {
		score += results[round]
	}
	return
}
