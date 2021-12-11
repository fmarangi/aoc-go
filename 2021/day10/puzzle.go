package day10

import (
	"sort"
	"strings"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func openStack(line string) (stack []rune, illegal rune) {
	for _, c := range line {
		switch c {
		case '(', '[', '{', '<':
			stack = append([]rune{c}, stack...)
		default:
			if abs(int(c-stack[0])) > 2 {
				return nil, c
			}
			stack = stack[1:]
		}
	}
	return stack, 0
}

func Part1(input string) (score int) {
	points := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	for _, line := range strings.Split(input, "\n") {
		_, illegal := openStack(line)
		if illegal != 0 {
			score += points[illegal]
		}
	}
	return
}

func Part2(input string) int {
	scores := []int{}
	points := map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}

	for _, line := range strings.Split(input, "\n") {
		stack, _ := openStack(line)
		if stack != nil {
			score := 0
			for i, n := 0, len(stack); i < n; i++ {
				score = score*5 + points[stack[i]]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}
