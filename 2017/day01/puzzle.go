package day01

import "strconv"

func Part1(input string) int {
	return CountMatches(input, 1)
}

func Part2(input string) int {
	return CountMatches(input, len(input)/2)
}

func CountMatches(input string, offset int) (match int) {
	n := len(input)
	for i, c := range input {
		if c == int32(input[(i+offset)%n]) {
			v, _ := strconv.ParseInt(string(c), 10, 0)
			match += int(v)
		}
	}
	return match
}
