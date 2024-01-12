package day14

import (
	"fmt"
	"strings"
)

type Point struct{ X, Y int }

func (a Point) Add(b Point) Point {
	return Point{a.X + b.X, a.Y + b.Y}
}

type Row [1000]bool
type Chart []Row

func (c Chart) Get(p Point) bool {
	return c[p.Y][p.X]
}

func (chart *Chart) PourSand() (bottom bool, firstStep bool) {
	var (
		max     = len(*chart)
		sand    = Point{500, 0}
		offsets = [3]Point{{0, 1}, {-1, 1}, {1, 1}}
		moving  = true
	)

	for moving && sand.Y < max {
		moving = false
		for _, o := range offsets {
			if p := sand.Add(o); p.Y >= max || !chart.Get(p) {
				moving, sand, firstStep = true, p, true
				break
			}
		}
	}

	if sand.Y < max {
		(*chart)[sand.Y][sand.X] = true
	}

	return sand.Y == max, firstStep
}

func Part1(input string) (units int) {
	chart := parseInput(input)
	for bottom := false; !bottom; units++ {
		bottom, _ = chart.PourSand()
	}
	return units - 1
}

func Part2(input string) (units int) {
	var (
		chart  = parseInput(input)
		bottom Row
	)

	for i := 0; i < len(bottom); i++ {
		bottom[i] = true
	}

	chart = append(chart, Row{}, bottom)
	for falling := true; falling; units++ {
		_, falling = chart.PourSand()
	}

	return
}

func parseInput(input string) Chart {
	var walls [][]Point
	for _, row := range strings.Split(input, "\n") {
		walls = append(walls, parseWall(row))
	}

	var max int
	for _, w := range walls {
		for _, p := range w {
			if p.Y > max {
				max = p.Y
			}
		}
	}

	var chart = make(Chart, max+1)
	for _, w := range walls {
		for i, n := 0, len(w); i < n-1; i++ {
			for _, p := range getPath(w[i], w[i+1]) {
				chart[p.Y][p.X] = true
			}
		}
	}

	return chart
}

func parseWall(input string) (wall []Point) {
	for _, p := range strings.Split(input, " -> ") {
		var pt Point
		fmt.Sscanf(p, "%d,%d", &pt.X, &pt.Y)
		wall = append(wall, pt)
	}
	return
}

func getPath(a, b Point) (path []Point) {
	var (
		step = Point{direction(a.X, b.X), direction(a.Y, b.Y)}
		curr = a
	)

	path = append(path, curr)
	for curr != b {
		curr.X += step.X
		curr.Y += step.Y
		path = append(path, curr)
	}

	return
}

func direction(a, b int) int {
	switch {
	case a < b:
		return 1
	case a > b:
		return -1
	}
	return 0
}
