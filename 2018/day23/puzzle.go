package day23

import (
	"fmt"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Nanobot struct {
	X, Y, Z, R int
}

func (a Nanobot) Distance(b Nanobot) int {
	return utils.Abs(a.X-b.X) + utils.Abs(a.Y-b.Y) + utils.Abs(a.Z-b.Z)
}

func (a Nanobot) InRange(b Nanobot) bool {
	return a.Distance(b) < b.R
}

func HighestFrequency(ranges [][]Nanobot) (highest []Nanobot) {
	max := 0
	for _, r := range ranges {
		if n := len(r); n > max {
			highest, max = r, n
		}
	}
	return
}

func Part1(input string) (inRange int) {
	var (
		bots    = ParseInput(input)
		largest Nanobot
	)

	for _, n := range bots {
		if n.R > largest.R {
			largest = n
		}
	}

	for _, n := range bots {
		if n.Distance(largest) <= largest.R {
			inRange++
		}
	}

	return
}

func Part2(input string) (distance int) {
	var (
		bots   = ParseInput(input)
		count  = len(bots)
		ranges = make([][]Nanobot, count)
		origin Nanobot
	)

	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if bots[i].InRange(bots[j]) {
				ranges[i] = append(ranges[i], bots[j])
			}
			if bots[j].InRange(bots[i]) {
				ranges[j] = append(ranges[j], bots[i])
			}
		}
	}

	for _, bot := range HighestFrequency(ranges) {
		if d := bot.Distance(origin) - bot.R; d > distance {
			distance = d
		}
	}

	return
}

func ParseInput(input string) (bots []Nanobot) {
	var x, y, z, r int
	for _, row := range strings.Split(input, "\n") {
		fmt.Sscanf(row, "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r)
		bots = append(bots, Nanobot{x, y, z, r})
	}
	return
}
