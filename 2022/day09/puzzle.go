package day09

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Point struct{ X, Y int }

func (a *Point) Add(b Point) {
	a.X += b.X
	a.Y += b.Y
}

func (a *Point) Follow(b Point) {
	x, y := b.X-a.X, b.Y-a.Y
	switch {
	case x/2 == 0 && y/2 != 0:
		a.Add(Point{x, y / 2})
	case y/2 == 0 && x/2 != 0:
		a.Add(Point{x / 2, y})
	default:
		a.Add(Point{x / 2, y / 2})
	}
}

func Head(moves string) <-chan Point {
	var (
		out     = make(chan Point)
		current = new(Point)
		offsets = map[string]Point{
			"U": {0, 1}, "D": {0, -1},
			"L": {-1, 0}, "R": {1, 0},
		}
	)

	go func() {
		defer close(out)
		for _, move := range strings.Split(moves, "\n") {
			parts := strings.Fields(move)
			for i, n := 0, utils.ToInt(parts[1]); i < n; i++ {
				current.Add(offsets[parts[0]])
				out <- *current
			}
		}
	}()

	return out
}

func TenKnotRope(head <-chan Point) <-chan Point {
	out := make(chan Point)

	// Init rope
	rope := [9]*Point{}
	for i := 0; i < 9; i++ {
		rope[i] = new(Point)
	}

	go func() {
		defer close(out)
		for h := range head {
			rope[0].Follow(h)
			for i := 1; i < 9; i++ {
				rope[i].Follow(*rope[i-1])
			}
			out <- *rope[8]
		}
	}()

	return out
}

func Part1(input string) int {
	var points = make(map[Point]bool)
	var tail = new(Point)
	for h := range Head(input) {
		tail.Follow(h)
		points[*tail] = true
	}
	return len(points)
}

func Part2(input string) int {
	var points = make(map[Point]bool)
	for p := range TenKnotRope(Head(input)) {
		points[p] = true
	}
	return len(points)
}
