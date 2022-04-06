package day11

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type position struct{ x, y int }

var offsets = map[string]position{
	"n":  {0, 2},
	"ne": {1, 1},
	"se": {1, -1},
	"s":  {0, -2},
	"sw": {-1, -1},
	"nw": {-1, 1},
}

func (a position) distance(b position) int {
	return (utils.Abs(a.x-b.x) + utils.Abs(a.y-b.y) + 1) / 2
}

func (from position) path(steps ...string) <-chan position {
	path, current := make(chan position), from
	go func() {
		defer close(path)
		for _, step := range steps {
			o := offsets[step]
			current = position{current.x + o.x, current.y + o.y}
			path <- current
		}
	}()
	return path
}

func Move(steps ...string) int {
	var start, current position
	for current = range start.path(steps...) {
	}
	return current.distance(start)
}

func Part1(input string) int {
	return Move(parseInput(input)...)
}

func Part2(input string) (max int) {
	var start, current position
	for current = range start.path(parseInput(input)...) {
		if d := current.distance(start); d > max {
			max = d
		}
	}
	return
}

func parseInput(input string) []string {
	return strings.Split(input, ",")
}
