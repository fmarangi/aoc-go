package day09

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func FindZero(history []int) (int, int) {
	var (
		head    = history[0]
		tail    = history[len(history)-1]
		allZero bool
		value   int
		factor  = -1
	)

	for !allZero {
		var (
			n    = len(history)
			next = make([]int, 0, n-1)
		)

		allZero = true
		for i := 1; i < n; i++ {
			value = history[i] - history[i-1]
			next = append(next, value)
			if value != 0 {
				allZero = false
			}
		}

		if !allZero {
			tail += value
			head += next[0] * factor
			factor *= -1
			history = next
		}
	}

	return head, tail
}

func Part1(input string) (sum int) {
	for _, history := range parseInput(input) {
		_, tail := FindZero(history)
		sum += tail
	}
	return
}

func Part2(input string) (sum int) {
	for _, history := range parseInput(input) {
		head, _ := FindZero(history)
		sum += head
	}
	return
}

func parseInput(input string) (histories [][]int) {
	for _, row := range strings.Split(input, "\n") {
		histories = append(histories, utils.Ints(strings.Fields(row)))
	}

	return
}
