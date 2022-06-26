package day03

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type move struct {
	dir   string
	steps int
}

type wire []move

type point struct{ x, y int }

func (p point) distance(o point) int {
	return utils.Abs(p.x-o.x) + utils.Abs(p.y-o.y)
}

func path(moves []move) (out chan point) {
	out = make(chan point)
	dirs := map[string]point{
		"U": {0, 1},
		"D": {0, -1},
		"R": {1, 0},
		"L": {-1, 0},
	}
	go func() {
		defer close(out)
		curr := point{}
		for _, m := range moves {
			for i := 0; i < m.steps; i++ {
				curr.x += dirs[m.dir].x
				curr.y += dirs[m.dir].y
				out <- curr
			}
		}
	}()
	return
}

func Part1(input string) int {
	var (
		wires = parseInput(input)
		seen  = map[point]bool{}
		cross = []int{}
	)

	for p := range path(wires[0]) {
		seen[p] = true
	}
	for p := range path(wires[1]) {
		if seen[p] {
			cross = append(cross, p.distance(point{}))
		}
	}
	return utils.Min(cross...)
}

func Part2(input string) int {
	var (
		wires = parseInput(input)
		seen  = map[point]int{}
		cross = []int{}
		i     = 0
	)

	for p := range path(wires[0]) {
		i++
		seen[p] = i
	}

	i = 0
	for p := range path(wires[1]) {
		i++
		if seen[p] > 0 {
			cross = append(cross, seen[p]+i)
		}
	}
	return utils.Min(cross...)
}

func parseInput(input string) (wires [2]wire) {
	for i, row := range strings.Split(input, "\n") {
		for _, m := range strings.Split(row, ",") {
			wires[i] = append(wires[i], move{m[0:1], utils.ToInt(m[1:])})
		}
	}
	return
}
