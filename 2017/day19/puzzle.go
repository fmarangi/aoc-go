package day19

import (
	"strings"
)

const (
	Up = iota
	Right
	Down
	Left
)

func followRoute(route string) (letters string, steps int) {
	max, start := len(route), strings.Index(route, "|")
	width := strings.Index(route, "\n") + 1
	offsets := [4]int{-width, 1, width, -1}

	for pos, dir := start, Down; route[pos] != ' '; pos += offsets[dir] {
		curr := route[pos]
		switch {
		case curr == '|', curr == '-':
			// Do nothing
		case curr == '+':
			for i := 0; i < 2; i++ {
				d := (dir + i*2 + 1) % 4
				if p := pos + offsets[d]; p <= max && p >= 0 && route[p] != ' ' {
					dir = d
					break
				}
			}
		default:
			letters += string(curr)
		}
		steps++
	}
	return
}

func Part1(input string) (letters string) {
	letters, _ = followRoute(input)
	return
}

func Part2(input string) (steps int) {
	_, steps = followRoute(input)
	return
}
