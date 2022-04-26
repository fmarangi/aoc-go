package day15

import (
	"fmt"
	"strings"
)

type Disc struct{ Pos, Num int }

func escape(discs []Disc, time int) bool {
	for _, disc := range discs {
		if (disc.Pos+time)%disc.Num != 0 {
			return false
		}
	}
	return true
}

func Part1(input string) (time int) {
	discs := parseInput(input)
	for !escape(discs, time) {
		time++
	}
	return
}

func Part2(input string) (time int) {
	discs := parseInput(input)
	discs = append(discs, Disc{len(discs) + 1, 11})
	for !escape(discs, time) {
		time++
	}
	return
}

func parseInput(input string) (discs []Disc) {
	var pattern = "Disc #%d has %d positions; at time=0, it is at position %d."
	for _, row := range strings.Split(input, "\n") {
		var d, n, p int
		fmt.Sscanf(row, pattern, &d, &n, &p)
		discs = append(discs, Disc{(p + d) % n, n})
	}
	return
}
