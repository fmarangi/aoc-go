package day12

import (
	"math"
	"strings"
)

type Route struct{ pos, steps int }

func canClimb(a, b byte) bool {
	if a == 'S' {
		a = 'a'
	}
	switch int8(b) - int8(a) {
	case 1, 0, -1, -2:
		return true
	}
	return false
}

func FindRoute(input string, start int) int {
	var (
		max   = len(input)
		width = strings.Index(input, "\n") + 1
		queue = []Route{{start, 0}}
		seen  = map[int]bool{start: true}
		curr  Route
	)

	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		for _, n := range [4]int{1, -1, width, -width} {
			pos := curr.pos + n
			switch {
			case pos < 0, pos >= max, seen[pos]:
				continue
			case input[pos] == 'E' && input[curr.pos] >= 'y':
				return curr.steps + 1
			case canClimb(input[curr.pos], input[pos]):
				queue = append(queue, Route{pos, curr.steps + 1})
				seen[pos] = true
			}
		}
	}
	return -1
}

func findB(input string) (b []int) {
	for i, c := range input {
		if c == 'b' {
			b = append(b, i)
		}
	}
	return
}

func Part1(input string) int {
	return FindRoute(input, strings.IndexByte(input, 'S'))
}

func Part2(input string) (min int) {
	min = math.MaxInt32
	for _, b := range findB(input) {
		if p := FindRoute(input, b) + 1; p > 0 && p < min {
			min = p
		}
	}
	return
}
